package model

import (
	"time"
)

// Interface
type DocumentCommentInterface interface {
	GetById(int64, int64)										(DocumentComment, error)
	Create(int64, int64, string, time.Time)	(int64, error)
	QueryAll(userId, documentId int64)      ([]DocumentComment, error)
	Update()																(error)
	Delete()																(int64, error)
	DeleteById(int64)												(int64, error)
}

// Request model
type AddDocumentCommentParams struct {
	PostMemberId	int64		`json:"postMemberId" validate:"required"`
	Content				string	`json:"content" validate:"required"`
}

type UpdateDocumentCommentParams struct {
	PostMemberId	int64		`json:"postMemberId" validate:"required"`
	Content				string	`json:"content" validate:"required"`
}

// Database model
type DocumentComment struct {
	Id							int64			`json:"id"`
	DocumentId 			int64			`json:"documentId"`
	DocumentName		string		`json:"documentName"`
	PostMemberId 		int64			`json:"postMemberId"`
	PostMemberName 	string		`json:"postMemberName"`
	Content					string		`json:"content"`
	CreatedAt 			time.Time	`json:"createdAt"`
}

// Method
func (DC *DocumentComment) GetById(documentId, documentCommentId int64) (DocumentComment, error) {
	sqlStatement := `
		SELECT 
		dc.id, dc.post_member_id, dc.content, dc.created_at,
		d.id as d_id, d.name as d_name,
		m.id as m_id, m.name as m_name   
		FROM document_comments dc
		INNER JOIN documents d 
		ON d.id = dc.document_id 
		INNER JOIN members  m
		ON m.id = dc.post_member_id 
		WHERE dc.id = $1 and dc.document_id = $2;`

	row := DbConf.PgConn.SQL.QueryRow(sqlStatement, documentCommentId, documentId)
	err := row.Scan(&DC.Id, &DC.PostMemberId, &DC.Content, &DC.CreatedAt, &DC.DocumentId, &DC.DocumentName, &DC.PostMemberId, &DC.PostMemberName)
	if err != nil {
		return DocumentComment{}, err
	}

	return *DC, nil
}

func (DC *DocumentComment) Create(documentId, postMemberId int64, content string, createdAt time.Time) (int64, error) {
	sqlStatement := `
		INSERT INTO document_comments
		(document_id, post_member_id, content, created_at)
		VALUES ($1, $2, $3, $4) RETURNING id;`

	var id int64
	err := DbConf.PgConn.SQL.QueryRow(sqlStatement, documentId, postMemberId, content, createdAt).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (DC *DocumentComment) QueryAll(userId, documentId int64) ([]DocumentComment, error) {
	sqlStatement := `
		SELECT
		dc.id, dc.post_member_id, dc.content, dc.created_at,
		d.id as d_id, d.name as d_name,
		m.id as m_id, m.name as m_name
		FROM document_comments dc
		INNER JOIN documents d
		ON d.id = dc.document_id
		INNER JOIN members  m
		ON m.id = dc.post_member_id
		WHERE dc.document_id = $1
		ORDER BY dc.id DESC;`

	rows, err := DbConf.PgConn.SQL.Query(sqlStatement, documentId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var documentComments []DocumentComment
	for rows.Next() {
		var documentComment DocumentComment
		err := rows.Scan(&documentComment.Id, &documentComment.PostMemberId, &documentComment.Content, &documentComment.CreatedAt, &documentComment.DocumentId, &documentComment.DocumentName, &documentComment.PostMemberId, &documentComment.PostMemberName)
		if err != nil {
			return nil, err
		}

		documentComments = append(documentComments, documentComment)
	}

	return documentComments, nil
}

func (DC *DocumentComment) Update() (error) {
	sqlStatement := `
		UPDATE document_comments
		SET post_member_id = $1, content = $2
		WHERE id = $3;`

	_, err := DbConf.PgConn.SQL.Exec(sqlStatement, DC.PostMemberId, DC.Content, DC.Id)
	if err != nil {
		return err
	}

	return nil
}

func (DC *DocumentComment) Delete() (int64, error) {
	sqlStatement := `
		DELETE FROM document_comments
		WHERE id = $1 RETURNING id;`

	var documentCommentId int64
	row := DbConf.PgConn.SQL.QueryRow(sqlStatement, DC.Id)
	err := row.Scan(&documentCommentId)
	if err != nil {
		return 0, err
	}

	return documentCommentId, nil
}

func (DC *DocumentComment) DeleteById(documentId int64) (int64, error) {
	sqlStatement := `
		DELETE FROM document_comments
		WHERE document_id = $1 RETURNING id;`

	var id int64
	row := DbConf.PgConn.SQL.QueryRow(sqlStatement, documentId)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}