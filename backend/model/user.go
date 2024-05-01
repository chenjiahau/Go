package model

type User struct {
	UserId		int64		`json:"id"`
	Email			string	`json:"email"`
	Username	string	`json:"username"`
	Password	string	`json:"password"`
	Token			string	`json:"token"`
	Expires		float64	`json:"expires"`
}