package model

import (
	"ivanfun.com/mis/internal/util"
)

// Interface
// For the methods that have to be implemented by the User struct
type UserInterface interface {
	GetId(SignUpParams)		int64
	Create(SignUpParams)	error
	Query(SignInParams)		error
}

// Request model
// Struct for the parameters of the sign up request
type SignUpParams struct {
	Email						string	`json:"email" validate:"required,email"`
	Name						string	`json:"username" validate:"required,min=1,max=32"`
	Password				string	`json:"password" validate:"required,min=8,max=20"`
	ConfirmPassword	string	`json:"confirmPassword" validate:"required,min=8,max=20"`
}

// Struct for the parameters of the sign in request
type SignInParams struct {
	Email			string	`json:"email" validate:"required,email"`
	Password	string	`json:"password" validate:"required,min=8,max=20"`
}

// Database model
// Struct for the user corresponding to the users table in the database
type User struct {
	Id				int64		`json:"id"`
	Email			string	`json:"email"`
	Name			string	`json:"username"`
	Password	string	`json:"password"`
	Token			string	`json:"token"`
}

// Methods
func NewUser() UserInterface {
	return &User{}
}

// Method to get the id of the user with the given email
func (U *User) GetId(sp SignUpParams) int64 {
	sqlStatement := `SELECT id FROM users WHERE email = $1`
	row := DbConf.PgConn.SQL.QueryRow(sqlStatement, sp.Email)

	var id int64
	err := row.Scan(&id)

	if err != nil {
		return 0
	}

	return id
}

func (U *User) Create(sp SignUpParams) error {
	sqlStatement := `INSERT INTO users (email, username, password) VALUES ($1, $2, $3);`
	_, err := DbConf.PgConn.SQL.Exec(sqlStatement, sp.Email, sp.Name, sp.Password)

	if err != nil {
		return err
	}

	return nil
}

func (U *User) Query(si SignInParams) error {
	sqlStatement := `SELECT id, email, username, password FROM users WHERE email= $1;`
	row := DbConf.PgConn.SQL.QueryRow(sqlStatement, si.Email)

	err := row.Scan(&U.Id, &U.Email, &U.Name, &U.Password)
	if err != nil {
		return err
	}

	err = util.CheckPasswordHash(si.Password, U.Password)
	if err != nil {
		return err
	}

	return nil
}
