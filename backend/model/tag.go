package model

import "fmt"

// Interface
type TagInterface interface {
	GetById(int64)												(TagDetail, error)
	GetByName(int64, string)											(Tag, error)
	Create(int64, string)									(int64, error)
	QueryAll()														([]TagDetail, error)
	QueryTotalCount(int64)								(int64, error)
	QueryByPage(int64, int, int, string, string)	([]TagDetail, error)
	Update(int64, int64, string)					(error)
	DeleteById(int64)											(Tag, error)
}

// Request model
type AddTagParams struct {
	ColorId	int64		`json:"colorId" validate:"required"`
	Name		string	`json:"name" validate:"required"`
}

type UpdateTagParams struct {
	ColorId	int64		`json:"colorId" validate:"required"`
	Name		string	`json:"name" validate:"required"`
}

// Database model
type Tag struct {
	Id			int64		`json:"id"`
	ColorId	int64		`json:"colorId"`
	Name		string	`json:"name"`
}

type TagDetail struct {
	Id								int64		`json:"id"`
	ColorCategoryId		int64		`json:"colorCategoryId"`
	ColorCategoryName string	`json:"colorCategoryName"`
	ColorId						int64		`json:"colorId"`
	ColorName 				string	`json:"colorName"`
	ColorHexCode			string	`json:"colorHexCode"`
	ColorRGBCode			string	`json:"colorRgbCode"`
	Name							string	`json:"name"`
}

// Method
func (C *Tag) GetById(id int64) (TagDetail, error) {
	sqlStatement := `
		SELECT t.id, cc.id, cc.name, c.id, c.name, c.hex_code, c.rgb_code, t.name
		FROM tags t
		INNER JOIN colors c
		ON t.color_id = c.id
		INNER JOIN color_categories cc 
		ON c.category_id = cc.id
		WHERE t.id = $1;`
	
	var tag TagDetail
	err := DbConf.PgConn.SQL.QueryRow(sqlStatement, id).Scan(&tag.Id, &tag.ColorCategoryId, &tag.ColorCategoryName, &tag.ColorId, &tag.ColorName, &tag.ColorHexCode, &tag.ColorRGBCode, &tag.Name)	
	if err != nil {
		return TagDetail{}, err
	}

	return tag, nil
}

func (C *Tag) GetByName(userId int64, name string) (Tag, error) {
	sqlStatement := `
	  SELECT id FROM tags WHERE name = $1
		WHERE id IN (SELECT tag_id FROM user_tags WHERE user_id = $2)
		;`
	
	var tag Tag
	err := DbConf.PgConn.SQL.QueryRow(sqlStatement, name, userId).Scan(&tag.Id)
	if err != nil {
		return Tag{}, err
	}

	return tag, nil
}

func (C *Tag) Create(colorId int64, name string) (int64, error) {
	sqlStatement := `INSERT INTO tags (color_id, name) VALUES ($1, $2) RETURNING id;`
	
	var id int64
	err := DbConf.PgConn.SQL.QueryRow(sqlStatement, colorId, name).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (C *Tag) QueryAll() ([]TagDetail, error) {
	sqlStatement := `
		SELECT t.id, cc.id, cc.name, c.id, c.name, c.hex_code, c.rgb_code, t.name
		FROM tags t
		INNER JOIN colors c
		ON t.color_id = c.id
		INNER JOIN color_categories cc 
		ON c.category_id = cc.id;`

	rows, err := DbConf.PgConn.SQL.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tags []TagDetail
	for rows.Next() {
		var tag TagDetail

		err := rows.Scan(&tag.Id, &tag.ColorCategoryId, &tag.ColorCategoryName, &tag.ColorId, &tag.ColorName, &tag.ColorHexCode, &tag.ColorRGBCode, &tag.Name)
		if err != nil {
			return nil, err
		}

		tags = append(tags, tag)
	}

	return tags, nil
}

func (C *Tag) QueryTotalCount(userId int64) (int64, error) {
	sqlStatement := `
	  SELECT COUNT(*) FROM tags
		WHERE id IN (SELECT tag_id FROM user_tags WHERE user_id = $1);`

	var count int64
	err := DbConf.PgConn.SQL.QueryRow(sqlStatement, userId).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (C *Tag) QueryByPage(userId int64, number, size int, orderBy, order string) ([]TagDetail, error) {
	switch orderBy {
	case "id":
		orderBy = "t.id"
	case "colorCategory":
		orderBy = "cc.id"
	case "color":
		orderBy = "c.name"
	case "name":
		orderBy = "t.name"
	default:
		orderBy = "id"
	}

	sqlStatement := fmt.Sprintf(`
	  SELECT t.id, cc.id, cc.name, c.id, c.name, c.hex_code, c.rgb_code, t.name
		FROM tags t
		INNER JOIN colors c
		ON t.color_id = c.id
		INNER JOIN color_categories cc 
		ON c.category_id = cc.id
		WHERE t.id IN (SELECT tag_id FROM user_tags WHERE user_id = %d)
		ORDER BY %s %s LIMIT $1 OFFSET $2;`,
		userId, orderBy, order)

	rows, err := DbConf.PgConn.SQL.Query(sqlStatement, size, (number - 1) * size)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tags []TagDetail
	for rows.Next() {
		var tag TagDetail
		err := rows.Scan(&tag.Id, &tag.ColorCategoryId, &tag.ColorCategoryName, &tag.ColorId, &tag.ColorName, &tag.ColorHexCode, &tag.ColorRGBCode, &tag.Name)

		if err != nil {
			return nil, err
		}

		tags = append(tags, tag)
	}

	return tags, nil
}

func (C *Tag) Update(id, colorId int64, name string) (error) {
	sqlStatement := `UPDATE tags SET color_id = $1, name = $2 WHERE id = $3;`

	_, err := DbConf.PgConn.SQL.Exec(sqlStatement, colorId, name, id)
	if err != nil {
		return err
	}

	return nil
}

func (C *Tag) DeleteAll() (error) {
	sqlStatement := `DELETE FROM tags;`
	_, err := DbConf.PgConn.SQL.Exec(sqlStatement)
	if err != nil {
		return err
	}

	return nil
}

func (C *Tag) DeleteById(id int64) (Tag, error) {
	sqlStatement := `DELETE FROM tags WHERE id = $1 RETURNING id;`

	var tag Tag
	err := DbConf.PgConn.SQL.QueryRow(sqlStatement, id).Scan(&tag.Id)
	if err != nil {
		return Tag{}, err
	}

	return tag, nil
}