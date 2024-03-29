package router

import (
	"net/http"

	"example.com/project/page"
)

func Router() {
	http.HandleFunc("/", page.Index)
}