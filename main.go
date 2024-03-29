package main

import (
	"log"
	"net/http"

	"example.com/project/router"
)

const portNumber = ":8080"

func main() {
	router.Router()

	log.Println("Starting application on port", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
