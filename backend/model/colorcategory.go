package model

// Interface
type ColorCategoryInterface interface {
	Create(string)	(int64, error)
	QueryAll()			([]ColorCategory, error)
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

func (C *ColorCategory) QueryAll() ([]ColorCategory, error) {
	sqlStatement := `SELECT id, name FROM color_categories;`

	rows, err := DbConf.PgConn.SQL.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var colorCategories []ColorCategory
	for rows.Next() {
		var colorCategory ColorCategory

		err := rows.Scan(&colorCategory.Id, &colorCategory.Name)
		if err != nil {
			return nil, err
		}

		colorCategories = append(colorCategories, colorCategory)
	}

	return colorCategories, nil
}

func (C *ColorCategory) DeleteAll() (error) {
	sqlStatement := `DELETE FROM color_categories;`
	_, err := DbConf.PgConn.SQL.Exec(sqlStatement)
	if err != nil {
		return err
	}

	return nil
}