package page

import (
	"encoding/json"
	"net/http"

	"example.com/project/data"
	"example.com/project/util"
)

type ResponseTokenData struct {
	Email 			string `json:"email"`
	Token 			string `json:"token"`
	ExpiredTime int64 `json:"expiredTime"`
}

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

var users = map[string]User{
	"john": {
		Email:    "john@test.com", 
		Password: "password123",
	},
	"jane": {
		Email:    "jane@test.com",
		Password: "password456",
	},
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user User
  err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	hasUser := false
	for _, v := range users {
		if v.Email == user.Email && v.Password == user.Password {
			hasUser = true
			break
		}
	}

	if !hasUser {
		rd := data.ResponseData{
			Success: false,
			Message: "Invalid email or password",
		}

		json.NewEncoder(w).Encode(rd)
		return
	}

	tokenString, expiredTime, err := util.CreateToken(user.Email)

	if err != nil {
		rd := data.ResponseData{
			Success: false,
			Message: "Failed to create token",
		}

		json.NewEncoder(w).Encode(rd)
		return
	}

	rd := ResponseTokenData{
		Email: user.Email,
		Token: tokenString,
		ExpiredTime: expiredTime,
	}

	json.NewEncoder(w).Encode(rd)
}