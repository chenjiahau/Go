package router

import (
	"fmt"
	"net/http"

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