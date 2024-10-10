package router

import (
	"net/http"
	"strings"

	"ivanfun.com/mis/internal/handler"
	"ivanfun.com/mis/internal/model"
	"ivanfun.com/mis/internal/util"
)

// func WriteToConsole(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Println("Executing route", r.URL.Path)
// 		logMessage := fmt.Sprintf("Executing route %s", r.URL.Path)
// 		util.WriteInfoLog(logMessage)
// 		next.ServeHTTP(w, r)
// 	})
// }

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
					Token: tokenString,
				})
			}
		}

		next.ServeHTTP(w, r)
	})
}

func CheckTokenAlive(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		res := handler.CheckTokenAlive()

		if !res {
			resErr := map[string]interface{}{
				"code": http.StatusUnauthorized,
				"message": "Unauthorized access",
			}

			util.ResponseJSONWriter(w, http.StatusUnauthorized, util.GetResponse(nil, resErr))
			return
		}

		next.ServeHTTP(w, r)
	})
}