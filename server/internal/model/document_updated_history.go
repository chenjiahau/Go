package model

import (
	"time"
)

// Interface
type DocumentUpdatedHistoryInterface interface {
	Create(int64, time.Time)	(int64, error)
}

// Request model
type AddDocumentUpdatedHistoryParams struct {
	DocumentId	int64	`json:"documentId"`
}

// Database model
type DocumentUpdatedHistory struct {
	Id					int64			`json:"id"`
	DocumentId	int64			`json:"documentId"`
	CreatedAt		time.Time	`json:"createdAt"`
}

// Method
func NewDocumentUpdatedHistory() DocumentUpdatedHistoryInterface {
	return &DocumentUpdatedHistory{}
}

func (duh *DocumentUpdatedHistory) Create(documentId int64, createdAt time.Time) (int64, error) {
	sqlStatement := `
		INSERT INTO document_updated_histories
		(document_id, created_at)
		VALUES ($1, $2) RETURNING id;`

	var id int64
	err := DbConf.PgConn.SQL.QueryRow(sqlStatement, documentId, createdAt).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}