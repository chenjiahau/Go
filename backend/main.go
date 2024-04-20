package main

import (
	"log"
	"net/http"

	"ivanfun.com/mis/config"
	"ivanfun.com/mis/handler"
	"ivanfun.com/mis/router"
)

func main() {
	hc := handler.NewConfig()
	handler.NewHandler(hc)

	srv := &http.Server{
		Addr:    config.Addr,
		Handler: router.GetRoutes(),
	}

	log.Printf("Server is running on %s\n", config.Addr)

	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
