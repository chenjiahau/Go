package page

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"example.com/project/data"
	"example.com/project/util"
	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type CreateUserParams struct {
	Username string `json:"username" validate:"required,min=1,max=32"`
	Password string `json:"password" validate:"required,min=8,max=20"`
	ConfirmPassword string `json:"confirmPassword" validate:"required,min=8,max=20"`
	FullName string `json:"fullName,omitempty" validate:"omitempty,min=1,max=64"`
}

func (repo *Repository) SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var p CreateUserParams
	err := json.NewDecoder(r.Body).Decode(&p)

	if err != nil {
		rd := data.ResponseData{
			Success: false,
			Message: "Bad request",
		}

		json.NewEncoder(w).Encode(rd)
		return
	}

	validate = validator.New()
	err = validate.Struct(p)

	if err != nil {
		rd := data.ResponseData{
			Success: false,
			Message: "Invalid data",
		}

		json.NewEncoder(w).Encode(rd)
		return
	}

	if p.Password != p.ConfirmPassword {
		rd := data.ResponseData{
			Success: false,
			Message: "Passwords do not match",
		}

		json.NewEncoder(w).Encode(rd)
		return
	}

	sqlStatement := `SELECT id FROM users WHERE username = $1`
	row := repo.DBConfig.PgConn.SQL.QueryRow(sqlStatement, p.Username)

	var id int
	err = row.Scan(&id)

	if err != sql.ErrNoRows {
		rd := data.ResponseData{
			Success: false,
			Message: "Username already exists",
		}

		json.NewEncoder(w).Encode(rd)
		return
	}

	password, err := bcrypt.GenerateFromPassword([]byte(p.Password), 14)

	if err != nil {
		rd := data.ResponseData{
			Success: false,
			Message: "Error creating user",
		}

		json.NewEncoder(w).Encode(rd)
		return
	}

	sqlStatement = `INSERT INTO users (username, password, full_name) VALUES ($1, $2, $3) RETURNING id`
	_, err = repo.DBConfig.PgConn.SQL.Exec(sqlStatement, p.Username, password, p.FullName)

	if err != nil {
		rd := data.ResponseData{
			Success: false,
			Message: "Error creating user",
		}

		json.NewEncoder(w).Encode(rd)
		return
	}

	rd := data.ResponseData{
		Success: true,
		Message: "User created successfully",
	}

	json.NewEncoder(w).Encode(rd)
}

type LoginParams struct {
	Username string `json:"username" validate:"required,min=1,max=32"`
	Password string `json:"password" validate:"required,min=8,max=20"`
}

type LoginResponse struct {
	Success bool `json:"success"`
	Message string `json:"message"`
	Data 	  map[string]interface{} `json:"data"`
}

func (repo *Repository) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var p LoginParams
	err := json.NewDecoder(r.Body).Decode(&p)

	if err != nil {
		rd := data.ResponseData{
			Success: false,
			Message: "Bad request",
		}

		json.NewEncoder(w).Encode(rd)
		return
	}

	validate = validator.New()
	err = validate.Struct(p)

	if err != nil {
		rd := data.ResponseData{
			Success: false,
			Message: "Invalid data",
		}

		json.NewEncoder(w).Encode(rd)
		return
	}

	sqlStatement := `SELECT id, password FROM users WHERE username = $1`
	row := repo.DBConfig.PgConn.SQL.QueryRow(sqlStatement, p.Username)

	var userId int
	var password string
	err = row.Scan(&userId, &password)

	if err == sql.ErrNoRows {
		rd := data.ResponseData{
			Success: false,
			Message: "Invalid username or password",
		}

		json.NewEncoder(w).Encode(rd)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(password), []byte(p.Password))
	if err != nil {
		rd := data.ResponseData{
			Success: false,
			Message: "Invalid username or password",
		}

		json.NewEncoder(w).Encode(rd)
		return
	}

	tokenString, expiredTime, err := util.CreateToken(p.Username)
	if err != nil {
		rd := data.ResponseData{
			Success: false,
			Message: "Error creating token",
		}

		json.NewEncoder(w).Encode(rd)
		return
	}

	sqlStatement = `INSERT INTO user_tokens (user_id, token, expired_at) VALUES ($1, $2, $3)`
	_, err = repo.DBConfig.PgConn.SQL.Exec(sqlStatement, userId, tokenString, time.Unix(expiredTime, 0))

	if err != nil {
		rd := data.ResponseData{
			Success: false,
			Message: "Error creating token",
		}

		json.NewEncoder(w).Encode(rd)
		return
	}

	rd := LoginResponse {
		Success: true,
		Message: "Login successful",
		Data: map[string]interface{}{
			"user_id": userId,
			"username": p.Username,
			"token": tokenString,
			"expired_at": time.Unix(expiredTime, 0),
		},
	}

	json.NewEncoder(w).Encode(rd)
}

func (repo *Repository) TestAuth(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		pd := data.ResponseData{
			Success: false,
			Message: "Unauthorized",
		}

		json.NewEncoder(w).Encode(pd)
		return
	}

	tokenString = tokenString[len("Bearer "):]
	err := util.VerifyToken(tokenString)
	if err != nil {
		pd := data.ResponseData{
			Success: false,
			Message: "Unauthorized",
		}

		json.NewEncoder(w).Encode(pd)
		return
	}

	pd := data.ResponseData{
		Success: true,
		Message: "Authorized",
	}

	json.NewEncoder(w).Encode(pd)
}