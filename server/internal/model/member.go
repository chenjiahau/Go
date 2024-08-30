package model

import "fmt"

// Interface
type MemberInterface interface {
	GetById(int64)																(Member, error)
	GetByName(int64, string)											(int64)
	Create(int64, string, bool)										(int64, error)
	QueryAll(int64)																([]Member, error)
	QueryTotalCount(int64)							  				(int64, error)
	QueryByPage(int64, int, int, string, string)	([]Member, error)
	Update(int64)																	(error)
	Delete(int64)																	(Member, error)
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
	Id							int64		`json:"id"`
	MemberRoleId		int64		`json:"memberRoleId"`
	MemberRoleTitle string	`json:"memberRoleTitle"`
	MemberRoleAbbr 	string	`json:"memberRoleAbbr"`
	Name 						string	`json:"name"`
	IsAlive					bool		`json:"isAlive"`
}

// Method
func NewMember() MemberInterface {
	return &Member{}
}

func (M *Member) GetById(id int64) (Member, error) {
	sqlStatement := `
		SELECT mr.id, mr.title, mr.abbr, m.id, m.name, m.is_alive
		FROM members m
		INNER JOIN member_roles mr
		ON m.member_role_id = mr.id
		WHERE m.id = $1;`

	var member Member
	row := DbConf.PgConn.SQL.QueryRow(sqlStatement, id)
	err := row.Scan(&member.MemberRoleId, &member.MemberRoleTitle, &member.MemberRoleAbbr, &member.Id, &member.Name, &member.IsAlive)
	if err != nil {
		return Member{}, err
	}

	return member, nil
}

func (M *Member) GetByName(userId int64, name string) (int64) {
	sqlStatement := `
		SELECT m.id
		FROM members m
		INNER JOIN member_roles mr
		ON m.member_role_id = mr.id
		WHERE m.name = $1 and m.id in (SELECT id FROM user_members where user_id = $2);`

	var member Member
	row := DbConf.PgConn.SQL.QueryRow(sqlStatement, name, userId)
	err := row.Scan(&member.Id)
	if err != nil {
		return 0
	}

	return member.Id
}

func (M *Member) Create(memberRoleId int64, name string, isAlive bool) (int64, error) {
	sqlStatement := `INSERT INTO members (member_role_id, name, is_alive) VALUES ($1, $2, $3) RETURNING id;`
	
	var id int64
	row := DbConf.PgConn.SQL.QueryRow(sqlStatement, memberRoleId, name, isAlive)
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (M *Member) QueryAll(userId int64) ([]Member, error) {
	sqlStatement := `
		SELECT mr.id, m.id, m.name , m.is_alive
		FROM members m
		INNER JOIN member_roles mr
		ON m.member_role_id = mr.id
		WHERE m.id IN (SELECT member_id FROM user_members WHERE user_id = $1);`

	rows, err := DbConf.PgConn.SQL.Query(sqlStatement, userId)
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

func (M *Member) QueryTotalCount(userId int64) (int64, error) {
	sqlStatement := `
	  SELECT COUNT(*) FROM members
		WHERE id in (SELECT member_id FROM user_members WHERE user_id = $1);`

	var count int64
	row := DbConf.PgConn.SQL.QueryRow(sqlStatement, userId)
	err := row.Scan(&count)
	if err != nil {
		return 0, err
	}

	return count, nil
}

func (M *Member) QueryByPage(userId int64, number, size int, orderBy, order string) ([]Member, error) {
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
		WHERE m.id IN (SELECT member_id FROM user_members WHERE user_id = %d)
		ORDER BY %s %s LIMIT $1 OFFSET $2;`,
		userId, orderBy, order)

	rows, err := DbConf.PgConn.SQL.Query(sqlStatement, size, (number - 1) * size)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var members []Member
	for rows.Next() {
		var member Member

		err := rows.Scan(&member.MemberRoleId, &member.MemberRoleTitle, &member.MemberRoleAbbr, &member.Id, &member.Name, &member.IsAlive)
		if err != nil {
			return nil, err
		}

		members = append(members, member)
	}

	return members, nil
}

func (M *Member) Update(userId int64) (error) {
	sqlStatement := `
		UPDATE members
		SET member_role_id = $1, name = $2, is_alive = $3
		WHERE id = $4 AND id IN (SELECT member_id FROM user_members WHERE user_id = $5);`

	_, err := DbConf.PgConn.SQL.Exec(sqlStatement, M.MemberRoleId, M.Name, M.IsAlive, M.Id, userId)
	if err != nil {
		return err
	}

	return nil
}

func (M *Member) Delete(userId int64) (Member, error) {
	sqlStatement := `
		DELETE FROM members
		WHERE id = $1
		AND id IN (SELECT member_id FROM user_members WHERE user_id = $2)
		AND id NOT IN (SELECT post_member_id FROM documents)
		AND id NOT IN (SELECT post_member_id FROM document_comments)
		RETURNING id;`

	var member Member
	row := DbConf.PgConn.SQL.QueryRow(sqlStatement, M.Id, userId)
	err := row.Scan(&member.Id)
	if err != nil {
		return Member{}, err
	}

	return member, nil
}