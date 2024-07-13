package model

import "fmt"

// Interface
type MemberInterface interface {
	GetById(int64)																(MemberDetail, error)
	GetByName(int64, string)											(Member, error)
	Create(int64, string, bool)										(int64, error)
	QueryAll()																		([]Member, error)
	QueryTotalCount(int64)							  				(int64, error)
	QueryByPage(int64, int, int, string, string)	([]MemberDetail, error)
	Update(int64, int64, string, bool)						(error)
	DeleteById(int64)															(Member, error)
}

// Request model
type AddMemberParams struct {
	MemberRoleId	int64		`json:"memberRoleId" validate:"required"`
	Name					string	`json:"name" validate:"required"`
	IsAlive				bool		`json:"isAlive" default:"true"`
}

type UpdateMemberParams struct {
	MemberRoleId	int64		`json:"memberRoleId" validate:"required"`
	Name					string	`json:"name" validate:"required"`
	IsAlive				bool		`json:"isAlive" default:"true"`
}

// Database model
type Member struct {
	Id						int64		`json:"id"`
	MemberRoleId	int64		`json:"memberRoleId"`
	Name					string	`json:"name"`
	IsAlive				bool		`json:"isAlive"`
}

type MemberDetail struct {
	Id							int64		`json:"id"`
	MemberRoleId		int64		`json:"memberRoleId"`
	MemberRoleTitle string	`json:"memberRoleTitle"`
	MemberRoleAbbr 	string	`json:"memberRoleAbbr"`
	Name 						string	`json:"name"`
	IsAlive					bool		`json:"isAlive"`
}

// Method
func (M *Member) GetById(id int64) (MemberDetail, error) {
	sqlStatement := `
		SELECT mr.id, mr.title, mr.abbr, t.id, t.name, t.is_alive
		FROM members t
		INNER JOIN member_roles mr
		ON t.member_role_id = mr.id
		WHERE t.id = $1;`

	var member MemberDetail
	err := DbConf.PgConn.SQL.QueryRow(sqlStatement, id).Scan(&member.MemberRoleId, &member.MemberRoleTitle, &member.MemberRoleAbbr, &member.Id, &member.Name, &member.IsAlive)
	if err != nil {
		return MemberDetail{}, err
	}

	return member, nil
}

func (M *Member) GetByName(userId int64, name string) (Member, error) {
	sqlStatement := `
	  SELECT id, member_role_id, name, is_alive FROM members
		WHERE name = $1 AND id in (SELECT member_id FROM user_members WHERE user_id = $2);`

	var member Member
	err := DbConf.PgConn.SQL.QueryRow(sqlStatement, name, userId).Scan(&member.Id, &member.MemberRoleId, &member.Name, &member.IsAlive)
	if err != nil {
		return Member{}, err
	}

	return member, nil
}

func (M *Member) Create(memberRoleId int64, name string, isAlive bool) (int64, error) {
	sqlStatement := `INSERT INTO members (member_role_id, name, is_alive) VALUES ($1, $2, $3) RETURNING id;`

	
	var id int64
	err := DbConf.PgConn.SQL.QueryRow(sqlStatement, memberRoleId, name, isAlive).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (M *Member) QueryAll() ([]Member, error) {
	sqlStatement := `
		SELECT cc.id, m.id, m.name , m.is_alive
		FROM members m
		INNER JOIN member_roles cc
		ON m.member_role_id = cc.id;`

	rows, err := DbConf.PgConn.SQL.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var members []Member
	for rows.Next() {
		var member Member

		err := rows.Scan(&member.MemberRoleId, &member.Id, &member.Name, &member.IsAlive)
		if err != nil {
			return nil, err
		}

		members = append(members, member)
	}

	return members, nil
}

func (C *Member) QueryTotalCount(userId int64) (int64, error) {
	sqlStatement := `
	  SELECT COUNT(*) FROM members
		WHERE id in (SELECT member_id FROM user_members WHERE user_id=$1);`

	var count int64
	err := DbConf.PgConn.SQL.QueryRow(sqlStatement, userId).Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (C *Member) QueryByPage(userId int64, number, size int, orderBy, order string) ([]MemberDetail, error) {
	switch orderBy {
	case "id":
		orderBy = "m.id"
	case "memberRole":
		orderBy = "mr.id"
	case "status":
		orderBy = "m.is_alive"
	default:
		orderBy = "m.id"
	}

	sqlStatement := fmt.Sprintf(`
	  SELECT mr.id, mr.title, mr.abbr, m.id, m.name, m.is_alive
		FROM members m
		INNER JOIN member_roles mr
		ON m.member_role_id = mr.id
		WHERE m.id IN (SELECT member_id FROM user_members WHERE user_id=%d)
		ORDER BY %s %s LIMIT $1 OFFSET $2;`,
		userId, orderBy, order)

	rows, err := DbConf.PgConn.SQL.Query(sqlStatement, size, (number - 1) * size)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var members []MemberDetail
	for rows.Next() {
		var member MemberDetail

		err := rows.Scan(&member.MemberRoleId, &member.MemberRoleTitle, &member.MemberRoleAbbr, &member.Id, &member.Name, &member.IsAlive)
		if err != nil {
			return nil, err
		}

		members = append(members, member)
	}

	return members, nil
}

func (C *Member) Update(id, memberRoleId int64, name string, isAlive bool) (error) {
	sqlStatement := `UPDATE members SET member_role_id = $1, name = $2, is_alive = $3 WHERE id = $4;`

	_, err := DbConf.PgConn.SQL.Exec(sqlStatement, memberRoleId, name, isAlive, id)
	if err != nil {
		return err
	}

	return nil
}

func (M *Member) DeleteById(id int64) (Member, error) {
	sqlStatement := `DELETE FROM members WHERE id = $1 RETURNING id;`

	var member Member
	err := DbConf.PgConn.SQL.QueryRow(sqlStatement, id).Scan(&member.Id)
	if err != nil {
		return Member{}, err
	}

	return member, nil
}