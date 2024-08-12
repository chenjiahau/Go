package util

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)


type TokenClaims struct {
	*jwt.StandardClaims
	UserId		int64		`json:"userId"`
	UserName	string	`json:"userName"`
}

var secret = []byte("can-you-keep-a-secret?")

func CreateToken(userId int64, userName string) (string, error) {
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	exp := time.Now().Add(time.Hour * 24)
	token.Claims = &TokenClaims{
		&jwt.StandardClaims{
			ExpiresAt: exp.Unix(),
		},
		userId,
		userName,
	}

	val, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return val, nil
}

func GetClaimsFromToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}

	return nil, err
}

func IsTokenStillAlive(expires int64) bool {
	return time.Now().Unix() < expires
}