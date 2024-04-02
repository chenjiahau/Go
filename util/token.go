package util

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte("jwt-token-secret-key")

func CreateToken(email string) (string, int64, error) {
	expiredTime := time.Now().Add(time.Hour * 24).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, 
		jwt.MapClaims{ 
			"email": email, 
			"exp": expiredTime, 
		})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", 0, err
	}

	return tokenString, expiredTime, nil
}

func VerifyToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		 return secretKey, nil
	})
 
	if err != nil {
		 return err
	}
 
	if !token.Valid {
		 return fmt.Errorf("invalid token")
	}
 
	return nil
}