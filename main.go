package main

import (
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {
	getRouter()

	log.Println("Starting application on port", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
