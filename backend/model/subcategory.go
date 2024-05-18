package model

import (
	"time"
)

// Interface
type SubCategoryInterface interface {
	GetById(int64, int64)										(SubCategory, error)
	GetByName(int64, string)								(SubCategory, error)
	Create(int64, string, time.Time, bool)	(int64, error)
	QueryAll(int64)													([]SubCategory, error)
	Update(int64, string, bool)							(error)
	DeleteById(int64)												(SubCategory, error)
}

// Request model
type AddSubCategoryParams struct {
	Name				string	`json:"name" validate:"required,min=1,max=32"`
	IsAlive			bool		`json:"isAlive" default:"true"`
}

type UpdateSubCategoryParams struct {
	Name				string	`json:"name" validate:"required,min=1,max=32"`
	IsAlive			bool		`json:"isAlive" default:"true"`
}

// Database model
type SubCategory struct {
	Id					int64			`json:"id"`
	CategoryId	int64			`json:"categoryId"`
	Name				string		`json:"name"`
	CreatedAt		time.Time	`json:"createdAt"`
	IsAlive			bool			`json:"isAlive"`
}

// Method
func (C *SubCategory) GetById(id , subId int64) (SubCategory, error) {
	sqlStatement := `SELECT * FROM subcategories WHERE id = $2 and category_id = $1;`

	row := DbConf.PgConn.SQL.QueryRow(sqlStatement, id, subId)
	var subCategory SubCategory
	err := row.Scan(&subCategory.Id, &subCategory.CategoryId, &subCategory.Name, &subCategory.CreatedAt, &subCategory.IsAlive)
	if err != nil {
		return SubCategory{}, err
	}

	return subCategory, nil
}

func (C *SubCategory) GetByName(categoryId int64, name string) (SubCategory, error) {
	sqlStatement := `SELECT * FROM subcategories WHERE category_id = $1 AND name = $2;`

	row := DbConf.PgConn.SQL.QueryRow(sqlStatement, categoryId, name)
	var subCategory SubCategory
	err := row.Scan(&subCategory.Id, &subCategory.CategoryId, &subCategory.Name, &subCategory.CreatedAt, &subCategory.IsAlive)
	if err != nil {
		return SubCategory{}, err
	}

	return subCategory, nil
}

func (C *SubCategory) Create(categoryId int64, name string, createdAt time.Time, isAlive bool) (int64, error) {
	sqlStatement := `INSERT INTO subcategories (category_id, name, created_at, is_alive) VALUES ($1, $2, $3, $4) RETURNING id;`

	var id int64
	err := DbConf.PgConn.SQL.QueryRow(sqlStatement, categoryId, name, createdAt, isAlive).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (C *SubCategory) QueryAll(categoryId int64) ([]SubCategory, error) {
	sqlStatement := `SELECT category_id, name, created_at, is_alive FROM subcategories WHERE category_id = $1;`

	rows, err := DbConf.PgConn.SQL.Query(sqlStatement, categoryId)
	if err != nil {
		return nil, err
	}

	var subCategories []SubCategory
	for rows.Next() {
		var subCategory SubCategory
		err := rows.Scan(&subCategory.CategoryId, &subCategory.Name, &subCategory.CreatedAt, &subCategory.IsAlive)
		if err != nil {
			return nil, err
		}

		subCategories = append(subCategories, subCategory)
	}

	return subCategories, nil
}

func (C *SubCategory) Update(id int64, name string, isAlive bool) error {
	sqlStatement := `UPDATE subcategories SET name = $2, is_alive = $3 WHERE id = $1;`

	_, err := DbConf.PgConn.SQL.Exec(sqlStatement, id, name, isAlive)
	if err != nil {
		return err
	}

	return nil
}

func (C *SubCategory) DeleteById(id int64) (SubCategory, error) {
	sqlStatement := `DELETE FROM subcategories WHERE id = $1 RETURNING id;`

	var subCategory SubCategory
	err := DbConf.PgConn.SQL.QueryRow(sqlStatement, id).Scan(&subCategory.Id)
	if err != nil {
		return SubCategory{}, err
	}

	return subCategory, nil
}