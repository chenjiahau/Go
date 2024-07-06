package model

// Interface
type MemberRoleInterface interface {
	Create(string, string)	(int64, error)
	QueryAll()							([]MemberRole, error)
	DeleteAll()							(error)
}

// Database model
type MemberRole struct {
	Id		int64		`json:"id"`
	Title	string	`json:"title"`
	Abbr  string	`json:"abbr"`
}

// Method
func (MR *MemberRole) Create(title, abbr string) (int64, error) {
	sqlStatement := `INSERT INTO member_roles (title, abbr) VALUES ($1, $2) RETURNING id;`

	var id int64
	err := DbConf.PgConn.SQL.QueryRow(sqlStatement, title, abbr).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (MR *MemberRole) QueryAll() ([]MemberRole, error) {
	sqlStatement := `SELECT id, title, abbr FROM member_roles;`

	rows, err := DbConf.PgConn.SQL.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var memberRoles []MemberRole
	for rows.Next() {
		var memberRole MemberRole

		err := rows.Scan(&memberRole.Id, &memberRole.Title, &memberRole.Abbr)
		if err != nil {
			return nil, err
		}

		memberRoles = append(memberRoles, memberRole)
	}

	return memberRoles, nil
}

func (MR *MemberRole) DeleteAll() (error) {
	sqlStatement := `DELETE FROM member_roles;`
	_, err := DbConf.PgConn.SQL.Exec(sqlStatement)
	if err != nil {
		return err
	}

	return nil
}