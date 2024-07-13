package model

// Interface
type UserTagInterface interface {
	Create(int64, int64)	(int64, error)
	DeleteById(int64)			(UserTag, error)
}

// Request model

// Database model
type UserTag struct {
	Id			int64	`json:"id"`
	UserId	int64	`json:"userId"`
	TagId		int64	`json:"tagId"`
}

// Method
func (C *UserTag) Create(userId, tagId int64) (int64, error) {
	sqlStatement := `INSERT INTO user_tags (user_id, tag_id) VALUES ($1, $2) RETURNING id;`

	var id int64
	err := DbConf.PgConn.SQL.QueryRow(sqlStatement, userId, tagId).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (C *UserTag) DeleteById(tagId int64) (UserTag, error) {
	sqlStatement := `DELETE FROM user_tags WHERE tag_id = $1 RETURNING id;`

	var userTag UserTag
	err := DbConf.PgConn.SQL.QueryRow(sqlStatement, tagId).Scan(&userTag.Id)

	if err != nil {
		return UserTag{}, err
	}

	return userTag, nil
}