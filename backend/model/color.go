package model

// Interface
type ColorInterface interface {
	Create(int64, string, string, string)	(int64, error)
	DeleteAll()														(error)
}

// Database model
type Color struct {
	Id				int64		`json:"id"`
	ColorName	string	`json:"color_name"`
	HexCode		string	`json:"hex_code"`
	RGBCode		string	`json:"rgb_code"`
}

// Method
func (C *Color) Create(categoryId int64, name, hexCode, rgbCode string) (int64, error) {
	sqlStatement := `INSERT INTO colors (category_id, color_name, hex_code, rgb_code) VALUES ($1, $2, $3, $4) RETURNING id;`
	
	var id int64
	err := DbConf.PgConn.SQL.QueryRow(sqlStatement, categoryId, name, hexCode, rgbCode).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (C *Color) DeleteAll() (error) {
	sqlStatement := `DELETE FROM colors;`
	_, err := DbConf.PgConn.SQL.Exec(sqlStatement)
	if err != nil {
		return err
	}

	return nil
}