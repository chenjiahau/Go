package model

import (
	"fmt"
	"time"
)

// Interface
type CategoryInterface interface {
	GetById(int64)												(CategorySimply, error)
	GetByName(string)											(CategorySimply, error)
	Create(string, time.Time, bool)				(int64, error)
	QueryAll()														([]Category, error)
	QueryTotalCount()											(int, error)
	QueryByPage(int, int, string, string)	([]CategorySimply, error)
	Update(int64, string, bool)						(error)
	DeleteById(int64)											(Category, error)
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
	Id					int64					`json:"id"`
	Name				string				`json:"name"`
	CreatedAt		time.Time			`json:"createdAt"`
	IsAlive			bool					`json:"isAlive"`
	SubCategory	[]SubCategory `json:"subcategories"`
}

type CategorySimply struct {
	Id								int64					`json:"id"`
	Name							string				`json:"name"`
	CreatedAt					time.Time			`json:"createdAt"`
	IsAlive						bool					`json:"isAlive"`
	SubCategoryCount	int64					`json:"subcategoryCount"`
}

// Method
func (C *Category) GetById(id int64) (CategorySimply, error) {
	sqlStatement := `
		SELECT c.id, c.name, c.created_at, c.is_alive,
		(SELECT COUNT(*) FROM subcategories sc WHERE sc.category_id  = c.id) AS subcategory_count
		FROM categories c WHERE c.id = $1;`

	row := DbConf.PgConn.SQL.QueryRow(sqlStatement, id)
	var category CategorySimply
	err := row.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.IsAlive, &category.SubCategoryCount)
	if err != nil {
		return CategorySimply{}, err
	}

	return category, nil
}

func (C *Category) GetByName(name string) (CategorySimply, error) {
	sqlStatement := `
		SELECT c.id, c.name, c.created_at, c.is_alive,
		(SELECT COUNT(*) FROM subcategories sc WHERE sc.category_id  = c.id) AS subcategory_count
		FROM categories c
		WHERE c.name = $1;`

	var category CategorySimply
	err := DbConf.PgConn.SQL.QueryRow(sqlStatement, name).Scan(&category.Id, &category.Name, &category.CreatedAt, &category.IsAlive, &category.SubCategoryCount)
	if err != nil {
		return CategorySimply{}, err
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

func (C *Category) QueryAll() ([]Category, error) {
	sqlStatement := `SELECT id, name, created_at, is_alive FROM categories;`

	rows, err := DbConf.PgConn.SQL.Query(sqlStatement)
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

		category.SubCategory = subCategories
		categories = append(categories, category)
	}

	return categories, nil
}

func (C *Category) QueryTotalCount() (int, error) {
	sqlStatement := `SELECT COUNT(*) FROM categories;`

	var count int
	err := DbConf.PgConn.SQL.QueryRow(sqlStatement).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (C *Category) QueryByPage(number int, size int, orderBy string, order string) ([]CategorySimply, error) {
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
		FROM categories c ORDER BY %s %s LIMIT $1 OFFSET $2;`, 
		orderBy, order)

	rows, err := DbConf.PgConn.SQL.Query(sqlStatement, size, (number - 1) * size)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categories []CategorySimply
	for rows.Next() {
		var category CategorySimply
		err := rows.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.IsAlive, &category.SubCategoryCount)

		if err != nil {
			return nil, err
		}

		categories = append(categories, category)
	}

	return categories, nil
}

func (C *Category) Update(id int64, name string, isAlive bool) (error) {
	sqlStatement := `UPDATE categories SET name = $1, is_alive = $2 WHERE id = $3;`

	_, err := DbConf.PgConn.SQL.Exec(sqlStatement, name, isAlive, id)
	if err != nil {
		return err
	}

	return nil
}

func (C *Category) DeleteById(id int64) (Category, error) {
	sqlStatement := `DELETE FROM categories WHERE id = $1 RETURNING id;`

	var category Category
	err := DbConf.PgConn.SQL.QueryRow(sqlStatement, id).Scan(&category.Id)
	if err != nil {
		return Category{}, err
	}

	return category, nil
}