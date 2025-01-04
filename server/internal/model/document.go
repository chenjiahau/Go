package model

import (
	"fmt"
	"time"

	"ivanfun.com/mis/internal/config"
)

// Interface
type DocumentInterface interface {
	GetById(int64, int64)																											(Document, error)
	GetByName(int64, string)																									(int64)
	Create(string, int64, int64, int64, []int64, []int64, string, time.Time)	(int64, error)
	QueryAll(int64)																														([]Document, error)
	QueryTotalCount(int64)																										(int64, error)
	QueryByPage(int64, int, int, string, string)															([]Document, error)
	Update(int64, []int64, []int64)																						(error)
	Delete(int64)																															(Document, error)
	QueryBySearch(int64, string)																							([]Document, error)
}

// Request model
type AddDocumentParams struct {
	Name							string	`json:"name" validate:"required"`
	CategoryId				int64		`json:"categoryId" validate:"required"`
	SubCategoryId			int64		`json:"subCategoryId" validate:"required"`
	PostMemberId			int64		`json:"postMemberId" validate:"required"`
	RelationMemberIds []int64	`json:"relationMemberIds" validate:"required"`
	TagIds						[]int64	`json:"tagIds" validate:"required"`
	Content						string	`json:"content" validate:"required"`
}

type UpdateDocumentParams struct {
	Name							string	`json:"name" validate:"required"`
	CategoryId				int64		`json:"categoryId" validate:"required"`
	SubCategoryId			int64		`json:"subCategoryId" validate:"required"`
	PostMemberId			int64		`json:"postMemberId" validate:"required"`
	RelationMemberIds []int64	`json:"relationMemberIds" validate:"required"`
	TagIds						[]int64	`json:"tagIds" validate:"required"`
	Content						string	`json:"content" validate:"required"`
}

// Database model
type Document struct {
	Id							int64											`json:"id"`
	Name 						string										`json:"name"`
	Category 				Category									`json:"category"`
	SubCategory 		SubCategory								`json:"subCategory"`
	PostMember 			Member										`json:"postMember"`
	RelationMembers []DocumentRelationMember	`json:"relationMembers"`
	Tags 						[]DocumentTag							`json:"tags"`
	Content					string										`json:"content"`
	CreatedAt 			time.Time									`json:"createdAt"`
	UpdatedAt 			time.Time									`json:"updatedAt"`
}

// Method
func NewDocument() DocumentInterface {
	return &Document{}
}

func (d *Document) GetById(userId, id int64) (Document, error) {
	sqlStatement := `
		SELECT
		d.id, d.name, d.category_id, d.subcategory_id, d.post_member_id, d.content, d.created_at,
		COALESCE((SELECT created_at FROM document_updated_histories WHERE document_id = d.id ORDER BY created_at DESC LIMIT 1), d.created_at) AS updated_at
		FROM documents d
		WHERE d.id = $1;`

	var name 					string
	var categoryId 		int64
	var subcategoryId int64
	var postMemberId 	int64
	var content 			string
	var createdAt 		time.Time
	var updatedAt 		time.Time

	row := DbConf.PgConn.SQL.QueryRow(sqlStatement, id)
	err := row.Scan(
		&id, &name, &categoryId, &subcategoryId, &postMemberId, &content,
		&createdAt, &updatedAt)
	if err != nil {
		return Document{}, err
	}

  var c CategoryInterface = &Category{}
	category, err := c.GetById(userId, categoryId)
	if err != nil {
		return Document{}, err
	}

	var sc SubCategoryInterface = &SubCategory{}
	subCategory, err := sc.GetById(categoryId, subcategoryId)
	if err != nil {
		return Document{}, err
	}

	var m MemberInterface = &Member{}
	postMember, err := m.GetById(postMemberId)
	if err != nil {
		return Document{}, err
	}

	var drm DocumentRelationMemberInterface = &DocumentRelationMember{}
	relationMembers, err := drm.GetById(id)
	if err != nil {
		return Document{}, err
	}

	var dt DocumentTagInterface = &DocumentTag{}
	tags, err := dt.GetByTags(id)
	if err != nil {
		return Document{}, err
	}

	return Document{
		Id: id,
		Name: name,
		Category: category,
		SubCategory: subCategory,
		PostMember: postMember,
		RelationMembers: relationMembers,
		Tags: tags,
		Content: content,
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
	}, nil
}

func (d *Document) GetByName(userId int64, name string) (int64) {
	sqlStatement := `
		SELECT d.id
		FROM documents d
		WHERE d.name = $1
		AND d.id IN (SELECT document_id FROM user_documents WHERE user_id = $2);`

	var documentId int64
	err := DbConf.PgConn.SQL.QueryRow(sqlStatement, name, userId).Scan(&documentId)
	if err != nil {
		return 0
	}

	return documentId
}

func (d *Document) Create(
	name string, categoryId, subcategoryId, postMemberId int64, 
	relationMemberIds, tagIds []int64, content string, createdAt time.Time) (int64, error) {
	sqlStatement := `
	  INSERT INTO documents
		(name, category_id, subcategory_id, post_member_id, content, created_at)
		VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;`

	var documentId int64
	var row = DbConf.PgConn.SQL.QueryRow(sqlStatement, name, categoryId, subcategoryId, postMemberId, content, createdAt)
	err := row.Scan(&documentId)
	if err != nil {
		return 0, err
	}

	// Insert relation members
	for _, relationMemberId := range relationMemberIds {
		sqlStatement = `INSERT INTO document_relation_members (document_id, member_id, created_at) VALUES ($1, $2, $3);`
		_, err = DbConf.PgConn.SQL.Exec(sqlStatement, documentId, relationMemberId, createdAt)
		if err != nil {
			return 0, err
		}
	}

	// Insert tags
	for _, tagId := range tagIds {
		sqlStatement = `INSERT INTO document_tags (document_id, tag_id, created_at) VALUES ($1, $2, $3);`
		_, err = DbConf.PgConn.SQL.Exec(sqlStatement, documentId, tagId, createdAt)
		if err != nil {
			return 0, err
		}
	}

	return documentId, nil
}

func (d *Document) QueryAll(userId int64) ([]Document, error) {
	sqlStatement := `
		SELECT d.id, d.name, d.category_id, d.subcategory_id, d.post_member_id, d.content, d.created_at
		FROM documents d ORDER BY d.created_at DESC;`;

	rows, err := DbConf.PgConn.SQL.Query(sqlStatement)
	if err != nil {
		return []Document{}, err
	}

	var documents []Document
	for rows.Next() {
		var id, categoryId, subcategoryId, postMemberId int64
		var name, content string
		var createdAt time.Time

		err = rows.Scan(&id, &name, &categoryId, &subcategoryId, &postMemberId, &content, &createdAt)
		if err != nil {
			return []Document{}, err
		}

		var c CategoryInterface = &Category{}
		category, err := c.GetById(userId, categoryId)
		if err != nil {
			return []Document{}, err
		}

		var sc SubCategoryInterface = &SubCategory{}
		subCategory, err := sc.GetById(categoryId, subcategoryId)
		if err != nil {
			return []Document{}, err
		}

		var m MemberInterface = &Member{}
		postMember, err := m.GetById(postMemberId)
		if err != nil {
			return []Document{}, err
		}

		var drm DocumentRelationMemberInterface = &DocumentRelationMember{}
		relationMembers, err := drm.GetById(id)
		if err != nil {
			return []Document{}, err
		}

		var dt DocumentTagInterface = &DocumentTag{}
		tags, err := dt.GetByTags(id)
		if err != nil {
			return []Document{}, err
		}

		documents = append(documents, Document{
			Id: id,
			Name: name,
			Category: category,
			SubCategory: subCategory,
			PostMember: postMember,
			RelationMembers: relationMembers,
			Tags: tags,
			Content: content,
			CreatedAt: createdAt,
		})
	}

	return documents, nil
}

func (d *Document) QueryTotalCount(userId int64) (int64, error) {
	sqlStatement := `
		SELECT COUNT(*) FROM documents
		WHERE id in (SELECT document_id FROM user_documents WHERE user_id=$1);`

		var count int64
		err := DbConf.PgConn.SQL.QueryRow(sqlStatement, userId).Scan(&count)
		if err != nil {
			return 0, err
		}

		return count, nil
}

func (d *Document) QueryByPage(userId int64, number, size int, orderBy, order string) ([]Document, error) {
	switch orderBy {
	case "id":
		orderBy = "id"
	case "name":
		orderBy = "name"
	case "category":
		orderBy = "category_name"
	case "subcategory":
		orderBy = "subcategory_name"
	case "post_member":
		orderBy = "post_member_name"
	case "created":
		orderBy = "created_at"
	case "updated":
		orderBy = "updated_at"
	default:
		orderBy = "updated_at"
	}

	sqlStatement := fmt.Sprintf(`
		SELECT
		d.id, d.name,
		d.category_id,
		(SELECT name FROM categories WHERE id = d.category_id) as category_name,
		d.subcategory_id,
		(SELECT name FROM subcategories s WHERE id = d.subcategory_id) as subcategory_name,
		d.post_member_id,
		(SELECT name FROM members s WHERE id = d.post_member_id) as post_member_name,
		d.content,
		d.created_at,
		COALESCE((SELECT created_at FROM document_updated_histories WHERE document_id = d.id ORDER BY created_at DESC LIMIT 1), d.created_at) AS updated_at
		FROM documents d
		WHERE d.id IN (SELECT document_id FROM user_documents WHERE user_id = %d)
		ORDER BY %s %s LIMIT $1 OFFSET $2;`,
		userId, orderBy, order)

	rows, err := DbConf.PgConn.SQL.Query(sqlStatement, size, (number-1)*size)
	if err != nil {
		return []Document{}, err
	}

	var documents []Document
	for rows.Next() {
		var id, categoryId, subcategoryId, postMemberId int64
		var name, categoryName, subCategoryName, postMemberName, content string
		var createdAt time.Time
		var updatedAt time.Time

		err = rows.Scan(
			&id, &name,
			&categoryId, &categoryName,
			&subcategoryId, &subCategoryName,
			&postMemberId, &postMemberName,
			&content,
			&createdAt, &updatedAt)

		if updatedAt.IsZero() {
			updatedAt = createdAt
		}

		if err != nil {
			return []Document{}, err
		}

		var c CategoryInterface = &Category{}
		category, err := c.GetById(userId, categoryId)
		if err != nil {
			return []Document{}, err
		}

		if !category.IsAlive {
			continue
		}

		var sc SubCategoryInterface = &SubCategory{}
		subCategory, err := sc.GetById(categoryId, subcategoryId)
		if err != nil {
			return []Document{}, err
		}

		if !subCategory.IsAlive {
			continue
		}

		var m MemberInterface = &Member{}
		postMember, err := m.GetById(postMemberId)
		if err != nil {
			return []Document{}, err
		}

		if !postMember.IsAlive {
			continue
		}

		var drm DocumentRelationMemberInterface = &DocumentRelationMember{}
		relationMembers, err := drm.GetById(id)
		if err != nil {
			return []Document{}, err
		}

		var dt DocumentTagInterface = &DocumentTag{}
		tags, err := dt.GetByTags(id)
		if err != nil {
			return []Document{}, err
		}

		documents = append(documents, Document{
			Id: id,
			Name: name,
			Category: category,
			SubCategory: subCategory,
			PostMember: postMember,
			RelationMembers: relationMembers,
			Tags: tags,
			Content: content,
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
		})
	}

	return documents, nil
}

func (d *Document) Update(userId int64, drmIds, tagIds []int64) (error) {
	sqlStatement := `
		UPDATE documents
		SET name = $1, category_id = $2, subcategory_id = $3, post_member_id = $4, content = $5
		WHERE id = $6 AND id IN (SELECT document_id FROM user_documents WHERE user_id = $7);`
	
	_, err := DbConf.PgConn.SQL.Exec(sqlStatement, d.Name, d.Category.Id, d.SubCategory.Id, d.PostMember.Id, d.Content, d.Id, userId)
	if err != nil {
		return err
	}

	// Delete relation members
	sqlStatement = `DELETE FROM document_relation_members WHERE document_id = $1;`
	_, err = DbConf.PgConn.SQL.Exec(sqlStatement, d.Id)
	if err != nil {
		return err
	}

	// Insert relation members
	for _, drmId := range drmIds {
		sqlStatement = `INSERT INTO document_relation_members (document_id, member_id, created_at) VALUES ($1, $2, $3);`
		_, err = DbConf.PgConn.SQL.Exec(sqlStatement, d.Id, drmId, time.Now())
		if err != nil {
			return err
		}
	}

	// Delete tags
	sqlStatement = `DELETE FROM document_tags WHERE document_id = $1;`
	_, err = DbConf.PgConn.SQL.Exec(sqlStatement, d.Id)
	if err != nil {
		return err
	}

	// Insert tags
	for _, tagId := range tagIds {
		sqlStatement = `INSERT INTO document_tags (document_id, tag_id, created_at) VALUES ($1, $2, $3);`
		_, err = DbConf.PgConn.SQL.Exec(sqlStatement, d.Id, tagId, time.Now())
		if err != nil {
			return err
		}
	}

	return nil
}

func (d *Document) Delete(userId int64) (Document, error) {
	sqlStatement := `
		DELETE FROM documents d
		WHERE d.id = $1 AND d.id IN (SELECT document_id FROM user_documents WHERE user_id = $2)
		RETURNING d.id, d.name, d.category_id, d.subcategory_id, d.post_member_id, d.content, d.created_at;`

	var document Document
	row := DbConf.PgConn.SQL.QueryRow(sqlStatement, d.Id, userId)
	err := row.Scan(&document.Id, &document.Name, &document.Category.Id, &document.SubCategory.Id,
		&document.PostMember.Id, &document.Content, &document.CreatedAt)
	if err != nil {
		return Document{}, err
	}

	return document, nil
}

func (d *Document) QueryBySearch(userId int64, keyword string) ([]Document, error) {
	sqlStatement := fmt.Sprintf(`
		SELECT
		DISTINCT d.id,
		d.name, d.category_id, d.subcategory_id, d.post_member_id, d.content, d.created_at
		FROM documents d
		INNER JOIN members m ON m.id  = d.post_member_id
		INNER JOIN categories c ON c.id = d.category_id
		INNER JOIN subcategories s ON s.id = d.subcategory_id
		FULL JOIN document_comments dc ON dc.document_id = d.id
		FULL JOIN document_tags dt ON dt.document_id = d.id
		WHERE
		d.id IN (SELECT document_id FROM user_documents WHERE user_id = %d)
		AND (UPPER(d.name) LIKE UPPER($1)
		OR UPPER(c.name) LIKE UPPER($1)
		OR UPPER(s.name) like UPPER($1)
		OR UPPER(m.name) LIKE UPPER($1)
		OR dt.tag_id IN (SELECT id FROM tags t WHERE t.id = dt.tag_id AND UPPER(t.name) LIKE UPPER($1))
		OR UPPER(d.content) LIKE UPPER($1)
		OR UPPER(dc.content) LIKE UPPER($1))
		ORDER BY d.created_at DESC
		LIMIT $2`, userId)

	rows, err := DbConf.PgConn.SQL.Query(sqlStatement, "%" + keyword + "%", config.App["SearchingLimit"].(int))
	if err != nil {
		return []Document{}, err
	}

	var documents []Document
	for rows.Next() {
		var id, categoryId, subcategoryId, postMemberId int64
		var name, content string
		var createdAt time.Time

		err = rows.Scan(&id, &name, &categoryId, &subcategoryId, &postMemberId, &content, &createdAt)
		if err != nil {
			return []Document{}, err
		}

		var c CategoryInterface = &Category{}
		category, err := c.GetById(userId, categoryId)
		if err != nil {
			return []Document{}, err
		}

		if !category.IsAlive {
			continue
		}

		var sc SubCategoryInterface = &SubCategory{}
		subCategory, err := sc.GetById(categoryId, subcategoryId)
		if err != nil {
			return []Document{}, err
		}

		if !subCategory.IsAlive {
			continue
		}

		var m MemberInterface = &Member{}
		postMember, err := m.GetById(postMemberId)
		if err != nil {
			return []Document{}, err
		}

		if !postMember.IsAlive {
			continue
		}

		var drm DocumentRelationMemberInterface = &DocumentRelationMember{}
		relationMembers, err := drm.GetById(id)
		if err != nil {
			return []Document{}, err
		}

		var dt DocumentTagInterface = &DocumentTag{}
		tags, err := dt.GetByTags(id)
		if err != nil {
			return []Document{}, err
		}

		documents = append(documents, Document{
			Id: id,
			Name: name,
			Category: category,
			SubCategory: subCategory,
			PostMember: postMember,
			RelationMembers: relationMembers,
			Tags: tags,
			Content: content,
			CreatedAt: createdAt,
		})
	}

	return documents, nil
}