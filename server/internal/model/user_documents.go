package model

// Interface
type UserDocumentInterface interface {
	Create(int64, int64)	(int64, error)
	DeleteById(int64)			()
	Delete(int64)					(UserDocument, error)
}

// Request model

// Database model
type UserDocument struct {
	Id					int64	`json:"id"`
	UserId			int64	`json:"userId"`
	DocumentId	int64	`json:"documentId"`
}

// Method
func NewUserDocument() UserDocumentInterface {
	return &UserDocument{}
}

func (ud *UserDocument) Create(userId, memberId int64) (int64, error) {
	sqlStatement := `INSERT INTO user_documents (user_id, document_id) VALUES ($1, $2) RETURNING id;`

	var id int64
	err := DbConf.PgConn.SQL.QueryRow(sqlStatement, userId, memberId).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (ud *UserDocument) DeleteById(memberId int64) () {
	sqlStatement := `DELETE FROM user_documents WHERE document_id = $1;`
	DbConf.PgConn.SQL.QueryRow(sqlStatement, memberId)
}

func (ud *UserDocument) Delete(documentId int64) (UserDocument, error) {
	sqlStatement := `DELETE FROM user_documents WHERE id = $1 RETURNING user_id, document_id;`

	var userDocument UserDocument
	err := DbConf.PgConn.SQL.QueryRow(sqlStatement, documentId).Scan(&userDocument.UserId, &userDocument.DocumentId)
	if err != nil {
		return UserDocument{}, err
	}

	return userDocument, nil
}