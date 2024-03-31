package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"example.com/project/config"
)

var tmpFuncs = template.FuncMap{}
var appConfig *config.AppConfig

func NewTemplates(ac *config.AppConfig) {
	appConfig = ac
}

func Render(w http.ResponseWriter, tl string) {
	var tc map[string]*template.Template

	if appConfig.UseCache {
		tc = appConfig.TemplateCache
	} else {
		tc, _ = CreateTemplateCache()
	}

	t, ok := tc[tl]
	if !ok {
		log.Fatal("Could not get template from template cache")
	}

	buf := new(bytes.Buffer)

	_ = t.Execute(buf, nil)

	_, err := buf.WriteTo(w)
	if err != nil {
		fmt.Println("error writing template to browser", err)
	}
}

func CreateTemplateCache() (map[string]*template.Template, error) {
	tmpCache := map[string]*template.Template{}

	pages, err := filepath.Glob("./templates/*.html")
	if err != nil {
		return tmpCache, err
	}

	for _, page := range pages {
		name := filepath.Base(page)
		ts, err := template.New(name).Funcs(tmpFuncs).ParseFiles(page)
		if err != nil {
			return tmpCache, err
		}

		matches, err := filepath.Glob("./templates/*.layout.html")
		if err != nil {
			return tmpCache, err
		}

		if len(matches) > 0 {
			ts, err = ts.ParseGlob("./templates/*.layout.html")
			if err != nil {
				return tmpCache, err
			}
		}

		tmpCache[name] = ts
	}

	return tmpCache, nil
}