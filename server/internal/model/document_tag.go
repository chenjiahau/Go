package model

import "time"

// Interface
type DocumentTagInterface interface {
	GetByTags(int64) 								([]DocumentTag, error)
	Create(int64, int64, time.Time)	(int64, error)
	Delete(int64)										(DocumentTag, error)
}

// Request model

// Database model
type DocumentTag struct {
	Id						int64			`json:"id"`
	DocumentId		int64			`json:"documentId"`
	TagId					int64			`json:"tagId"`
	TagName				string		`json:"tagName"`
	ColorId				int64			`json:"colorId"`
	ColorName			string		`json:"colorName"`
	ColorHexCode 	string		`json:"colorHexCode"`
	CreatedAt			time.Time	`json:"createdAt"`
}

// Method
func NewDocumentTag() DocumentTagInterface {
	return &DocumentTag{}
}

func (DT *DocumentTag) GetByTags(documentId int64) ([]DocumentTag, error) {
	sqlStatement := `
		SELECT dt.id, dt.document_id, dt.tag_id, dt.created_at,
		(
			SELECT c.id as color_id
			FROM tags t
			INNER JOIN colors c ON c.id = t.color_id 
			WHERE t.id = dt.tag_id
		),
		(
			SELECT c.name as color_name
			FROM tags t
			INNER JOIN colors c ON c.id = t.color_id 
			WHERE t.id = dt.tag_id
		),
		(
			SELECT c.hex_code as color_hex_code
			FROM tags t  
			INNER JOIN colors c ON c.id = t.color_id 
			WHERE t.id = dt.tag_id
		),
		(
			SELECT t.name as tag_name
			FROM tags t  
			WHERE t.id = dt.tag_id
		)
		FROM document_tags dt 
		WHERE document_id = $1;`

	rows, err := DbConf.PgConn.SQL.Query(sqlStatement, documentId)
	if err != nil {
		return []DocumentTag{}, err
	}

	var documentTags []DocumentTag
	for rows.Next() {
		var documentTag DocumentTag
		err = rows.Scan(
			&documentTag.Id, &documentTag.DocumentId, &documentTag.TagId, &documentTag.CreatedAt,
			&documentTag.ColorId, &documentTag.ColorName, &documentTag.ColorHexCode, &documentTag.TagName)

		if err != nil {
			return []DocumentTag{}, err
		}

		documentTags = append(documentTags, documentTag)
	}

	return documentTags, nil
}

func (DT *DocumentTag) Create(documentId, tagId int64, createdAt time.Time) (int64, error) {
	sqlStatement := `INSERT INTO document_tags (document_id, tag_id, created_at) VALUES ($1, $2, $3) RETURNING id;`

	var id int64
	err := DbConf.PgConn.SQL.QueryRow(sqlStatement, documentId, tagId, createdAt).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (DT *DocumentTag) Delete(tagId int64) (DocumentTag, error) {
	sqlStatement := `DELETE FROM document_tags WHERE id = $1 RETURNING document_id, tag_id, created_at;`

	var documentTag DocumentTag
	err := DbConf.PgConn.SQL.QueryRow(sqlStatement, tagId).Scan(&documentTag.DocumentId, &documentTag.TagId, &documentTag.CreatedAt)
	if err != nil {
		return DocumentTag{}, err
	}

	return documentTag, nil
}