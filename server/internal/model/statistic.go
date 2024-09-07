package model

// Interface
type MostPublisherInterface interface {
	Query(int64) ([]MostPublisher, error)
}

type MostCommentInterface interface {
	Query(int64) ([]MostComment, error)
}


// Database model
type MostPublisher struct {
	MemberId					int64		`json:"memberId"`
	MemberName				string	`json:"memberName"`
	NumberOfPost			int64		`json:"numberOfPost"`
}

type MostComment struct {
	DocumentId			int64 	`json:"documentId"`
	DocumentName		string 	`json:"documentName"`
	NumberOfComment int64 	`json:"numberOfComment"`
}

// Method
func (mp *MostPublisher) Query(userId int64) ([]MostPublisher, error) {
	sqlStatement := `
		SELECT
			m2.dpmi as member_id,
			(SELECT name FROM members WHERE id = m2.dpmi) as member_name,
			(m2.cdpmic + m2.cdcpmi) as number_of_post
		FROM
			(
			SELECt
				m1.dpmi,
				m1.cdpmic,
				(
				SELECT
					count(post_member_id)
				FROM
					document_comments dc
				WHERE
					post_member_id = m1.dpmi) AS cdcpmi
			FROM
				(
				SELECT
					d.post_member_id as dpmi,
					count(d.post_member_id) cdpmic
				FROM
					documents d
				WHERE d.id IN (SELECT document_id FROM user_documents ud WHERE user_id = $1)
				GROUP BY
					d.post_member_id) m1) m2
				ORDER BY number_of_post DESC
				LIMIT 10;`

	rows, err := DbConf.PgConn.SQL.Query(sqlStatement, userId)
	if err != nil {
		return nil, err
	}

	var mostPublishers []MostPublisher
	for rows.Next() {
		var mostPublisher MostPublisher

		err := rows.Scan(&mostPublisher.MemberId, &mostPublisher.MemberName, &mostPublisher.NumberOfPost)
		if err != nil {
			return nil, err
		}

		mostPublishers = append(mostPublishers, mostPublisher)
	}

	return mostPublishers, nil
}

func (mc *MostComment) Query(userId int64) ([]MostComment, error) {
	sqlStatement := `
		SELECT
			d.id AS document_id,
			d.name AS document_name,
			(SELECT count(*) FROM document_comments dc WHERE dc.document_id = d.id) AS number_of_comment
			FROM documents d
			WHERE d.id IN (SELECT document_id FROM user_documents ud WHERE user_id = $1)
		ORDER BY number_of_comment DESC
		LIMIT 10;`

	rows, err := DbConf.PgConn.SQL.Query(sqlStatement, userId)
	if err != nil {
		return nil, err
	}

	var mostComments []MostComment
	for rows.Next() {
		var mostComment MostComment

		err := rows.Scan(&mostComment.DocumentId, &mostComment.DocumentName, &mostComment.NumberOfComment)
		if err != nil {
			return nil, err
		}

		mostComments = append(mostComments, mostComment)
	}

	return mostComments, nil
}