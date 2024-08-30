package model

import "time"

// Interface
type DocumentRelationMemberInterface interface {
  GetById(int64)									([]DocumentRelationMember, error)
	Create(int64, int64, time.Time)	(int64, error)
	Delete(int64)										(DocumentRelationMember, error)
}

// Request model

// Database model
type DocumentRelationMember struct {
	Id							int64			`json:"id"`
	DocumentId			int64			`json:"documentId"`
	MemberRoleId		int64			`json:"memberRoleId"`
	MemberRoleTitle	string		`json:"memberRoleName"`
	MemberRoleAbbr	string		`json:"memberRoleAbbr"`
	MemberId				int64			`json:"memberId"`
	MemberName			string		`json:"name"`
	CreatedAt				time.Time	`json:"createdAt"`
}

// Method
func NewDocumentRelationMember() DocumentRelationMemberInterface {
	return &DocumentRelationMember{}
}

func (DRM *DocumentRelationMember) GetById(documentId int64) ([]DocumentRelationMember, error) {
	sqlStatement := `
	  SELECT drm.id, drm.document_id, drm.member_id, drm.created_at, m.name as member_name,
		(
			SELECT mr.id as member_role_id
			FROM member_roles mr
			WHERE mr.id = m.member_role_id
		),
		(
			SELECT mr.title as member_role_title
			FROM member_roles mr
			WHERE mr.id = m.member_role_id
		),
		(
			SELECT mr.abbr as member_role_abbr
			FROM member_roles mr
			WHERE mr.id = m.member_role_id
		)
		FROM document_relation_members drm
		INNER JOIN members m
		ON drm.member_id = m.id
		WHERE drm.document_id = $1;`

	rows, err := DbConf.PgConn.SQL.Query(sqlStatement, documentId)
	if err != nil {
		return []DocumentRelationMember{}, err
	}

	var documentRelationMembers []DocumentRelationMember
	for rows.Next() {
		var documentRelationMember DocumentRelationMember
		err = rows.Scan(
			&documentRelationMember.Id, &documentRelationMember.DocumentId, &documentRelationMember.MemberId,
			&documentRelationMember.CreatedAt, &documentRelationMember.MemberName,
			&documentRelationMember.MemberRoleId, &documentRelationMember.MemberRoleTitle, &documentRelationMember.MemberRoleAbbr)

		if err != nil {
			return []DocumentRelationMember{}, err
		}

		documentRelationMembers = append(documentRelationMembers, documentRelationMember)
	}

	return documentRelationMembers, nil
}

func (DRM *DocumentRelationMember) Create(documentId, memberId int64, createdAt time.Time) (int64, error) {
	sqlStatement := `INSERT INTO document_relation_members (document_id, member_id, created_at) VALUES ($1, $2, $3) RETURNING id;`

	var id int64
	err := DbConf.PgConn.SQL.QueryRow(sqlStatement, documentId, memberId, createdAt).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (DRM *DocumentRelationMember) Delete(drmId int64) (DocumentRelationMember, error) {
	sqlStatement := `DELETE FROM document_relation_members WHERE id = $1 RETURNING document_id, member_id, created_at;`

	var documentRelationMember DocumentRelationMember
	err := DbConf.PgConn.SQL.QueryRow(sqlStatement, drmId).Scan(&documentRelationMember.DocumentId, &documentRelationMember.MemberId, &documentRelationMember.CreatedAt)
	if err != nil {
		return DocumentRelationMember{}, err
	}

	return documentRelationMember, nil
}