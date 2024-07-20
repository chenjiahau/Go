package model

// Interface
type UserCategoryInterface interface {
	Create(int64, int64)	(int64, error)
	DeleteById(int64)			()
}

// Request model

// Database model
type UserCategory struct {
	Id					int64	`json:"id"`
	UserId			int64	`json:"userId"`
	CategoryId	int64	`json:"categoryId"`
}

// Method
func (UC *UserCategory) Create(userId, categoryId int64) (int64, error) {
	sqlStatement := `INSERT INTO user_categories (user_id, category_id) VALUES ($1, $2) RETURNING id;`

	var id int64
	err := DbConf.PgConn.SQL.QueryRow(sqlStatement, userId, categoryId).Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (UC *UserCategory) DeleteById(categoryId int64) () {
	sqlStatement := `DELETE FROM user_categories WHERE category_id = $1;`
	DbConf.PgConn.SQL.QueryRow(sqlStatement, categoryId)
}