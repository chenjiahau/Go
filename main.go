package main

import (
	"fmt"
	"log"
	"net/http"
)

const portNumber = ":8080"

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "This is the home page")
}

func main() {
	http.HandleFunc("/", Home)

	log.Println("Starting application on port", portNumber)
	_ = http.ListenAndServe(portNumber, nil)
}
