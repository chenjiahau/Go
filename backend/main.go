package main

import (
	"log"
	"net/http"

	"ivanfun.com/mis/config/server"
	"ivanfun.com/mis/router"
)

func main() {
	srv := &http.Server{
		Addr:    server.Addr,
		Handler: router.GetRoutes(),
	}

	log.Printf("Server is running on %s\n", server.Addr)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
