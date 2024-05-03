package router

import (
	"fmt"
	"net/http"
	"strings"

	"ivanfun.com/mis/handler"
	"ivanfun.com/mis/model"
	"ivanfun.com/mis/util"
)

func EnableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logMessage := fmt.Sprintf("Executing CORS middleware for route %s", r.URL.Path)
		util.WriteInfoLog(logMessage)

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PATCH, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers:", "Origin, Content-Type, X-Auth-Token, Authorization")
		w.Header().Set("Content-Type", "application/json")

		next.ServeHTTP(w, r)
 })
}

func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		logMessage := fmt.Sprintf("Executing route %s", r.URL.Path) 
		util.WriteInfoLog(logMessage)
		next.ServeHTTP(w, r)
	})
}

func ParseAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		handler.SetUser(nil)
		authHeader := r.Header.Get("Authorization")
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := util.GetClaimsFromToken(tokenString)

		if err != nil {
			handler.SetUser(nil)
		}

		if claims != nil {
			userId := int64(claims["userId"].(float64))
			userName := claims["userName"].(string)
			expires := claims["exp"].(float64)
			isTokenAlive := util.IsTokenStillAlive(int64(expires))

			if userId == 0 || userName == "" || !isTokenAlive {
				handler.SetUser(nil)
			} else {
				handler.SetUser(&model.User{
					Id: userId,
					Name: userName,
				})
			}
		}

		next.ServeHTTP(w, r)
	})
}