package router

import (
	"log"
	"net/http"
)

func WriteToConsole(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Executing route", r.URL.Path)
		next.ServeHTTP(w, r)
	})
}