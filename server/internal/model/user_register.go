package model

import "time"

// Interface
// For the methods that have to be implemented by the User register struct
type UserRegisterInterface interface {
	GetId(int64) 								(UserRegister, error)
	Create(UserRegisterParams) 	error
	Query(string) 							error
	Delete(int64) 							error
}

// Request model
// Struct for the parameters of the user register request
type UserRegisterParams struct {
	UserId    int64 		`json:"userId"`
	Token     string 		`json:"token"`
	ExpiredAt time.Time `json:"expiredAt"`
}

// Database model
// Struct for the user register corresponding to the user_registers table in the database
type UserRegister struct {
	Id        int64     `json:"id"`
	UserId    int64     `json:"userId"`
	Token     string    `json:"token"`
	ExpiredAt time.Time `json:"expiredAt"`
}

// Methods
// Method to create a new user register
func NewUserRegister() UserRegisterInterface {
	return &UserRegister{}
}

func (ur *UserRegister) GetId(int64) (UserRegister, error) {
	panic("unimplemented")
}

// Method to get the user register with the given user id
func (ur *UserRegister) Create(urp UserRegisterParams) (error) {
	sqlStatement := `INSERT INTO user_registers (user_id, token, expired_at) VALUES ($1, $2, $3) RETURNING id`
	row := DbConf.PgConn.SQL.QueryRow(sqlStatement, urp.UserId, urp.Token, urp.ExpiredAt)

	err := row.Scan(&ur.Id)
	if err != nil {
		return err
	}

	ur.UserId = urp.UserId
	ur.Token = urp.Token
	ur.ExpiredAt = urp.ExpiredAt

	return nil
}

// Method to query the user register with the token
func (ur *UserRegister) Query(token string) error {
	sqlStatement := `SELECT id, user_id, token, expired_at FROM user_registers WHERE token = $1`
	row := DbConf.PgConn.SQL.QueryRow(sqlStatement, token)

	err := row.Scan(&ur.Id, &ur.UserId, &ur.Token, &ur.ExpiredAt)
	if err != nil {
		return err
	}

	return nil
}

// Method to delete the user register with the given id
func (ur *UserRegister) Delete(id int64) error {
	sqlStatement := `DELETE FROM user_registers WHERE id = $1`
	_, err := DbConf.PgConn.SQL.Exec(sqlStatement, id)

	if err != nil {
		return err
	}

	return nil
}
