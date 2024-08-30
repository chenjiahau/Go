package model

import (
	"fmt"
	"time"
)

// Interface
type CategoryInterface interface {
	GetById(int64, int64)													(Category, error)
	GetByName(int64, string)											(Category, error)
	Create(string, time.Time, bool)								(int64, error)
	QueryAll(int64)																([]Category, error)
	QueryTotalCount(int64)												(int64, error)
	QueryByPage(int64, int, int, string, string)	([]Category, error)
	Update(int64)																	(error)
	Delete(int64)																			(Category, error)
}

// Request model
type AddCategoryParams struct {
	Name		string	`json:"name" validate:"required,min=1,max=32"`
	IsAlive bool		`json:"isAlive" default:"true"`
}

type UpdateCategoryParams struct {
	Name		string	`json:"name" validate:"required,min=1,max=32"`
	IsAlive bool		`json:"isAlive" default:"true"`
}

// Database model
type Category struct {
	Id								int64			`json:"id"`
	Name							string		`json:"name"`
	CreatedAt					time.Time	`json:"createdAt"`
	IsAlive						bool			`json:"isAlive"`
	SubCategoryCount	int64			`json:"subcategoryCount"`
}

// Method
func NewCategory() CategoryInterface {
	return &Category{}
}

func (C *Category) GetById(userId, id int64) (Category, error) {
	sqlStatement := `
		SELECT c.id, c.name, c.created_at, c.is_alive,
		(SELECT COUNT(*) FROM subcategories sc WHERE sc.category_id  = c.id) AS subcategory_count
		FROM categories c
		WHERE c.id = $1 AND c.id IN (SELECT category_id FROM user_categories WHERE user_id = $2);`

	var category Category
	row := DbConf.PgConn.SQL.QueryRow(sqlStatement, id, userId)
	err := row.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.IsAlive, &category.SubCategoryCount)
	if err != nil {
		return Category{}, err
	}

	return category, nil
}

func (C *Category) GetByName(userId int64, name string) (Category, error) {
	sqlStatement := `
		SELECT c.id, c.name, c.created_at, c.is_alive,
		(SELECT COUNT(*) FROM subcategories sc WHERE sc.category_id  = c.id) AS subcategory_count
		FROM categories c
		WHERE c.name = $1 AND c.id IN (SELECT category_id FROM user_categories WHERE user_id = $2);`

	var category Category
	row := DbConf.PgConn.SQL.QueryRow(sqlStatement, name, userId)
	err := row.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.IsAlive, &category.SubCategoryCount)
	if err != nil {
		return Category{}, err
	}

	return category, nil
}

func (C *Category) Create(name string, createdAt time.Time, isAlive bool) (int64, error) {
	sqlStatement := `INSERT INTO categories (name, created_at, is_alive) VALUES ($1, $2, $3) RETURNING id;`

	var id int64
	err := DbConf.PgConn.SQL.QueryRow(sqlStatement, name, createdAt, isAlive).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (C *Category) QueryAll(userId int64) ([]Category, error) {
	sqlStatement := `
		SELECT c.id, c.name, c.created_at, c.is_alive
		FROM categories c
		WHERE c.id IN (SELECT category_id FROM user_categories WHERE user_id = $1);`

	rows, err := DbConf.PgConn.SQL.Query(sqlStatement, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []Category
	for rows.Next() {
		var category Category

		err := rows.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.IsAlive)
		if err != nil {
			return nil, err
		}

		SubCategory := SubCategory{}
		subCategories, err := SubCategory.QueryAll(category.Id)
		if err != nil {
			return nil, err
		}

		category.SubCategoryCount = int64(len(subCategories))
		categories = append(categories, category)
	}

	return categories, nil
}

func (C *Category) QueryTotalCount(userId int64) (int64, error) {
	sqlStatement := `
		SELECT COUNT(*) FROM categories
		WHERE id in (SELECT category_id FROM user_categories WHERE user_id=$1);`

	var count int64
	err := DbConf.PgConn.SQL.QueryRow(sqlStatement, userId).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (C *Category) QueryByPage(userId int64, number, size int, orderBy, order string) ([]Category, error) {
	switch orderBy {
	case "id":
		orderBy = "id"
	case "name":
		orderBy = "name"
	case "created":
		orderBy = "created_at"
	case "subcategory":
		orderBy = "subcategory_count"
	case "status":
		orderBy = "is_alive"
	default:
		orderBy = "id"
	}

	sqlStatement := fmt.Sprintf(`
	  SELECT
		c.id, c.name, c.created_at, c.is_alive,
		(SELECT COUNT(*) FROM subcategories sc WHERE sc.category_id  = c.id) AS subcategory_count
		FROM categories c
		WHERE id IN (SELECT category_id FROM user_categories WHERE user_id=%d)
		ORDER BY %s %s LIMIT $1 OFFSET $2;`,
		userId, orderBy, order)

	rows, err := DbConf.PgConn.SQL.Query(sqlStatement, size, (number - 1) * size)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []Category
	for rows.Next() {
		var category Category
		err := rows.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.IsAlive, &category.SubCategoryCount)

		if err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}

func (C *Category) Update(userId int64) (error) {
	sqlStatement := `
		UPDATE categories
		SET name = $1, is_alive = $2
		WHERE id = $3 AND id IN (SELECT category_id FROM user_categories WHERE user_id = $4);`

	_, err := DbConf.PgConn.SQL.Exec(sqlStatement, C.Name, C.IsAlive, C.Id, userId)
	if err != nil {
		return err
	}

	return nil
}

func (C *Category) Delete(userId int64) (Category, error) {
	sqlStatement := `
		DELETE FROM categories
		WHERE id = $1
		AND id IN (SELECT category_id FROM user_categories WHERE user_id = $2)
		AND id NOT IN (SELECT category_id FROM documents)
		RETURNING id;`

	var category Category
	err := DbConf.PgConn.SQL.QueryRow(sqlStatement, C.Id, userId).Scan(&category.Id)
	if err != nil {
		return Category{}, err
	}

	return category, nil
}