package page

import (
	"net/http"

	"example.com/project/render"
)

func Index(w http.ResponseWriter, r *http.Request) {
	render.Render(w, "index.html")
}