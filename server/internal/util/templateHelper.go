package util

import (
	"text/template"
)

var tc = make(map[string]*template.Template)

func ParseTemplate(tmplPath string) (*template.Template, error) {
	var tmpl *template.Template
	var err error

	if tmpl, ok := tc[tmplPath]; ok {
		return tmpl, nil
	}

	tmpl, err = template.ParseFiles("template/"  + tmplPath + ".html")
	if err != nil {
		return nil, err
	}
	tc[tmplPath] = tmpl

	return tmpl, nil
}