package model

// Interface
type UserTagInterface interface {
	Create(int64, int64)	(int64, error)
	DeleteById(int64)			()
}

// Request model

// Database model
type UserTag struct {
	Id			int64	`json:"id"`
	UserId	int64	`json:"userId"`
	TagId		int64	`json:"tagId"`
}

// Method
func (UT *UserTag) Create(userId, tagId int64) (int64, error) {
	sqlStatement := `INSERT INTO user_tags (user_id, tag_id) VALUES ($1, $2) RETURNING id;`

	var id int64
	err := DbConf.PgConn.SQL.QueryRow(sqlStatement, userId, tagId).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (UT *UserTag) DeleteById(tagId int64) () {
	sqlStatement := `DELETE FROM user_tags WHERE tag_id = $1;`
	DbConf.PgConn.SQL.QueryRow(sqlStatement, tagId)
}