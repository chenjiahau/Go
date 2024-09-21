package model

import (
	"errors"

	"ivanfun.com/mis/internal/util"
)

// Interface
// For the methods that have to be implemented by the User struct
type UserInterface interface {
	GetId(SignUpParams)		int64
	Create(SignUpParams)	(int64, error)
	Query(SignInParams)		error
	Active(int64)					error
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
	Id						int64		`json:"id"`
	Email					string	`json:"email"`
	Name					string	`json:"username"`
	Password			string	`json:"password"`
	Token					string	`json:"token"`
	IsRegistered	bool		`json:"isRegistered"`
}

// Methods
func NewUser() UserInterface {
	return &User{}
}

// Method to get the id of the user with the given email
func (u *User) GetId(sp SignUpParams) int64 {
	sqlStatement := `SELECT id FROM users WHERE email = $1`
	row := DbConf.PgConn.SQL.QueryRow(sqlStatement, sp.Email)

	var id int64
	err := row.Scan(&id)

	if err != nil {
		return 0
	}

	return id
}

func (u *User) Create(sp SignUpParams) (int64, error) {
	sqlStatement := `INSERT INTO users (email, username, password) VALUES ($1, $2, $3) RETURNING id;`

	var userId int64
	err := DbConf.PgConn.SQL.QueryRow(sqlStatement, sp.Email, sp.Name, sp.Password).Scan(&userId)

	if err != nil {
		return 0, err
	}

	return userId, nil
}

func (u *User) Query(si SignInParams) error {
	sqlStatement := `SELECT id, email, username, password, is_registered FROM users WHERE email= $1;`
	row := DbConf.PgConn.SQL.QueryRow(sqlStatement, si.Email)

	err := row.Scan(&u.Id, &u.Email, &u.Name, &u.Password, &u.IsRegistered)
	if err != nil {
		return err
	}

	if !u.IsRegistered {
		return errors.New("User is not registered")
	}

	err = util.CheckPasswordHash(si.Password, u.Password)
	if err != nil {
		return err
	}

	return nil
}

func (u *User) Active(id int64) error {
	sqlStatement := `UPDATE users SET is_registered = TRUE WHERE id = $1;`
	_, err := DbConf.PgConn.SQL.Exec(sqlStatement, id)

	if err != nil {
		return err
	}

	return nil
}
