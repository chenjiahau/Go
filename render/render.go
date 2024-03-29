package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templateCache = make(map[string]*template.Template)

func createTemplateCache(t string) error {
	templates := []string {
		fmt.Sprintf("./templates/%s", t), "./templates/layout.html",
	}

	tpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}

	templateCache[t] = tpl

	return nil
}

func Render(w http.ResponseWriter, t string) {
	var tpl *template.Template
	var err error

	_, inMap := templateCache[t]
	if !inMap {
		log.Printf("Creating template from cache: %s\n", t)
		err = createTemplateCache(t)

		if err != nil {
			log.Println(err)
		}
	} else {
		log.Printf("Using template from cache: %s\n", t)
	}

	tpl = templateCache[t]

	err = tpl.Execute(w, nil)
	if err != nil {
		log.Println(err)
	}
}
