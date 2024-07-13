package model

// Interface
type UserMemberInterface interface {
	Create(int64, int64)	(int64, error)
	DeleteById(int64)			(UserMember, error)
}

// Request model

// Database model
type UserMember struct {
	Id				int64	`json:"id"`
	UserId		int64	`json:"userId"`
	MemberId	int64	`json:"memberId"`
}

// Method
func (C *UserMember) Create(userId, memberId int64) (int64, error) {
	sqlStatement := `INSERT INTO user_members (user_id, member_id) VALUES ($1, $2) RETURNING id;`

	var id int64
	err := DbConf.PgConn.SQL.QueryRow(sqlStatement, userId, memberId).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (C *UserMember) DeleteById(memberId int64) (UserMember, error) {
	sqlStatement := `DELETE FROM user_members WHERE member_id = $1 RETURNING id;`

	var userMember UserMember
	err := DbConf.PgConn.SQL.QueryRow(sqlStatement, memberId).Scan(&userMember.Id)

	if err != nil {
		return UserMember{}, err
	}

	return userMember, nil
}