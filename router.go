package main

import (
	"net/http"

	"example.com/project/page"
)

func getRouter() {
	http.HandleFunc("/", page.Index)
}