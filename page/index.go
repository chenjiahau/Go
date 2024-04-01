package page

import (
	"net/http"

	"example.com/project/data"
	"example.com/project/render"
)

func Index(w http.ResponseWriter, r *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "This is the index page."

	render.Render(w, "index.html", &data.TemplateData{
		StringMap: stringMap,
	}, r)
}