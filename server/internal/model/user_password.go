package model

import (
	"time"
)

// Interface
// For the methods that have to be implemented by the User reset password struct
type UserResetPasswordInterface interface {
	GetId(int64)										(UserResetPassword, error)
	Create(UserResetPasswordParams)	error
	Query(string)										(UserResetPassword, error)
	Invalidate(int64) 							error
	// Delete(int64)										error
}

// Request model
// Struct for the parameters of the user reset password
type UserResetPasswordParams struct {
	UserId		int64			`json:"userId"`
	Token			string		`json:"token"`
	ExpiredAt	time.Time	`json:"expiredAt"`
}

// Struct for the parameters of the change password request
type ChangePasswordParams struct {
	Email			string	`json:"email" validate:"required,email"`
	Token			string	`json:"token" validate:"required"`
	Password	string	`json:"password" validate:"required,min=8,max=20"`
}

// Database model
// Struct for the user register corresponding to the user_reset_passwords table in the database
type UserResetPassword struct {
	Id				int64			`json:"id"`
	UserId		int64			`json:"userId"`
	Token			string		`json:"token"`
	ExpiredAt	time.Time	`json:"expiredAt"`
	IsValid		bool			`json:"isValid"`
}

// Methods
// Method to create a new user register
func NewUserResetPassword() UserResetPasswordInterface {
	return &UserResetPassword{}
}

func (ur *UserResetPassword) GetId(int64) (UserResetPassword, error) {
	panic("unimplemented")
}

// Method to get the user reset password with the given user id
func (ur *UserResetPassword) Create(urp UserResetPasswordParams) (error) {
	sqlStatement := `INSERT INTO user_reset_passwords (user_id, token, expired_at, is_valid) VALUES ($1, $2, $3, true) RETURNING id;`
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

// Method to get the user reset password with the given email
func (ur *UserResetPassword) Query(email string) (UserResetPassword, error) {
	sqlStatement := `
		SELECT
		urp.id, urp.user_id, urp.token, urp.expired_at, urp.is_valid
		FROM user_reset_passwords urp
		INNER JOIN users u
		ON u.id = urp.user_id
		WHERE urp.is_valid = true AND u.email = $1
		ORDER BY urp.id desc
		LIMIT 1;`
	row := DbConf.PgConn.SQL.QueryRow(sqlStatement, email)

	var urp UserResetPassword
	err := row.Scan(&urp.Id, &urp.UserId, &urp.Token, &urp.ExpiredAt, &urp.IsValid)
	if err != nil {
		return UserResetPassword{}, err
	}

	return urp, nil
}

// Method to delete the user reset password with the given id
// func (ur *UserResetPassword) Delete(id int64) error {
// 	sqlStatement := `DELETE FROM user_reset_passwords WHERE user_id = $1;`
// 	_, err := DbConf.PgConn.SQL.Exec(sqlStatement, id)

// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// Method to invalidate the user reset password with the given id
func (ur *UserResetPassword) Invalidate(userId int64) error {
	sqlStatement := `UPDATE user_reset_passwords SET is_valid = false WHERE user_id = $1;`
	_, err := DbConf.PgConn.SQL.Exec(sqlStatement, userId)

	if err != nil {
		return err
	}

	return nil
}