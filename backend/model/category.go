package model

import (
	"time"
)

// Interface
type CategoryInterface interface {
	GetById(int64)									(Category, error)
	GetByName(string)								(Category, error)
	Create(string, time.Time, bool)	(int64, error)
	QueryAll()											([]Category, error)
	Update(int64, string, bool)			(error)
	DeleteById(int64)								(Category, error)
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
	SubCategory	[]SubCategory `json:"subCategory"`
}

// Method
func (C *Category) GetById(id int64) (Category, error) {
	sqlStatement := `SELECT id, name, created_at, is_alive FROM categories WHERE id = $1;`

	row := DbConf.PgConn.SQL.QueryRow(sqlStatement, id)
	var category Category
	err := row.Scan(&category.Id, &category.Name, &category.CreatedAt, &category.IsAlive)
	if err != nil {
		return Category{}, err
	}

	SubCategory := SubCategory{}
	subCategories, err := SubCategory.QueryAll(category.Id)
	if err != nil {
		return Category{}, err
	}

	category.SubCategory = subCategories

	return category, nil
}

func (C *Category) GetByName(name string) (Category, error) {
	sqlStatement := `SELECT * FROM categories WHERE name = $1;`

	var category Category
	err := DbConf.PgConn.SQL.QueryRow(sqlStatement, name).Scan(&category.Id, &category.Name, &category.CreatedAt, &category.IsAlive)
	if err != nil {
		return Category{}, err
	}

	subCategory := SubCategory{}
	subCategories, err := subCategory.QueryAll(category.Id)
	if err != nil {
		return Category{}, err
	}
	category.SubCategory = subCategories

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