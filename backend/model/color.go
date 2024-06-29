package model

// Interface
type ColorInterface interface {
	Create(int64, string, string, string)	(int64, error)
	QueryAll()														([]Color, error)
	DeleteAll()														(error)
}

// Database model
type Color struct {
	Id					int64		`json:"id"`
	CategoryId	int64		`json:"categoryId"`
	Name				string	`json:"name"`
	HexCode			string	`json:"hexCode"`
	RGBCode			string	`json:"rgbCode"`
}

// Method
func (C *Color) Create(categoryId int64, name, hexCode, rgbCode string) (int64, error) {
	sqlStatement := `INSERT INTO colors (category_id, name, hex_code, rgb_code) VALUES ($1, $2, $3, $4) RETURNING id;`

	
	var id int64
	err := DbConf.PgConn.SQL.QueryRow(sqlStatement, categoryId, name, hexCode, rgbCode).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (C *Color) QueryAll() ([]Color, error) {
	sqlStatement := `
		SELECT cc.id, c.id, c.name , c.hex_code , c.rgb_code
		FROM colors c
		INNER JOIN color_categories cc
		ON c.category_id = cc.id;`

	rows, err := DbConf.PgConn.SQL.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var colors []Color
	for rows.Next() {
		var color Color

		err := rows.Scan(&color.CategoryId, &color.Id, &color.Name, &color.HexCode, &color.RGBCode)
		if err != nil {
			return nil, err
		}

		colors = append(colors, color)
	}

	return colors, nil
}

func (C *Color) DeleteAll() (error) {
	sqlStatement := `DELETE FROM colors;`
	_, err := DbConf.PgConn.SQL.Exec(sqlStatement)
	if err != nil {
		return err
	}

	return nil
}