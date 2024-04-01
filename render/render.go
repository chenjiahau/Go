package render

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"path/filepath"

	"github.com/justinas/nosurf"

	"example.com/project/config"
	"example.com/project/data"
)

var tmpFuncs = template.FuncMap{}
var appConfig *config.AppConfig

func NewAppConfig(ac *config.AppConfig) {
	appConfig = ac
}

func AddDefaultData(td *data.TemplateData, r *http.Request) *data.TemplateData {
	td.CSRFToken = nosurf.Token(r)
	return td
}

func Render(w http.ResponseWriter, tl string, td *data.TemplateData, r *http.Request) {
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
	td = AddDefaultData(td, r)

	_ = t.Execute(buf, td)

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