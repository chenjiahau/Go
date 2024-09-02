package model

import (
	"time"
)

// Interface
// For the methods that have to be implemented by the User struct
type TokenInterface interface {
	GetByUserId(int64) 													([]Token, error)
	Create(int64, string, time.Time, time.Time)	error
	Query(string) 															(Token, error)
	SetIsAlive(string, bool)										error
}

// Database model
// Struct for the token corresponding to the tokens table in the database
type Token struct {
	UserId		int64			`json:"userId"`
	Token			string		`json:"token"`
	CreatedAt	time.Time	`json:"createdAt"`
	ExpiredAt	time.Time	`json:"expiredAt"`
	IsAlive		bool			`json:"isAlive"`
}

// Methods
func NewToken() TokenInterface {
	return &Token{}
}

// Method to get the token of the user with the given user id
func (t *Token) GetByUserId(userId int64) ([]Token, error) {
	sqlStatement := `SELECT token FROM tokens WHERE user_id = $1 and is_alive = true;`
	rows, err := DbConf.PgConn.SQL.Query(sqlStatement, userId)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tokens []Token
	for rows.Next() {
		var token Token

		err := rows.Scan(&token.Token)
		if err != nil {
			return nil, err
		}

		tokens = append(tokens, token)
	}

	return tokens, nil
}

func (t *Token) Create(userId int64, token string, createdAt time.Time, expiredAt time.Time) error {
	sqlStatement := `INSERT INTO tokens (user_id, token, created_at, expired_at, is_alive) VALUES ($1, $2, $3, $4, $5);`

	_, err := DbConf.PgConn.SQL.Exec(sqlStatement, userId, token, createdAt, expiredAt, true)
	if err != nil {
		return err
	}

	return nil
}

func (t *Token) Query(token string) (Token, error) {
	sqlStatement := `SELECT user_id, token, created_at, expired_at, is_alive FROM tokens WHERE token = $1;`
	row := DbConf.PgConn.SQL.QueryRow(sqlStatement, token)

	var rt Token
	err := row.Scan(&rt.UserId, &rt.Token, &rt.CreatedAt, &rt.ExpiredAt, &rt.IsAlive)
	if err != nil {
		return Token{}, err
	}

	return rt, nil
}

func (t *Token) SetIsAlive(token string, isAlive bool) error {
	sqlStatement := `UPDATE tokens SET is_alive = $1 WHERE token = $2;`

	_, err := DbConf.PgConn.SQL.Exec(sqlStatement, isAlive, token)
	if err != nil {
		return err
	}

	return nil
}