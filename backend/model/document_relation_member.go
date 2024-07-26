package model

// Interface
type DocumentRelationMemberInterface interface {
	// GetById(int64)																(Tag, error)
	// GetByName(int64, string)											(int64)
	Create(int64, string)													(int64, error)
	// QueryAll()																		([]Tag, error)
	// QueryTotalCount(int64)												(int64, error)
	// QueryByPage(int64, int, int, string, string)	([]Tag, error)
	// Update()																			(error)
	// Delete()																			(Tag, error)
}

// Request model

// Database model
type DocumentRelationMember struct {
}

// Method
