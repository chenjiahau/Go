package render

import (
	"fmt"
	"html/template"
	"net/http"
)

func Render(w http.ResponseWriter, tpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tpl, "./templates/layout.html")
	err := parsedTemplate.Execute(w, nil)

	if err != nil {
		fmt.Println("error parsing template:", err)
	}
}
