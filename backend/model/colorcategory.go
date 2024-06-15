package model

// Interface
type ColorCategoryInterface interface {
	Create(string)	(int64, error)
	DeleteAll()			(error)
}

// Database model
type ColorCategory struct {
	Id		int64		`json:"id"`
	Name	string	`json:"name"`
}

// Method
func (C *ColorCategory) Create(name string) (int64, error) {
	sqlStatement := `INSERT INTO color_categories (name) VALUES ($1) RETURNING id;`

	var id int64
	err := DbConf.PgConn.SQL.QueryRow(sqlStatement, name).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (C *ColorCategory) DeleteAll() (error) {
	sqlStatement := `DELETE FROM color_categories;`
	_, err := DbConf.PgConn.SQL.Exec(sqlStatement)
	if err != nil {
		return err
	}

	return nil
}