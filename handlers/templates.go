package handlers

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

// The layoutPath is where the boilerplate templates are kept.
// It is a directory.
const layoutPath = "templates/layout/"

// includePath contains all of the content templates
const includePath = "templates/"

var allTemplates = map[string]*template.Template{}

// Init loads the templates into Go's templater runtime.
func init() {
	var err error
	var tmplStr = `{{define "main" }} {{ template "base" . }} {{ end }}`

	layout, err := filepath.Glob(layoutPath + "*.tmpl")
	if err != nil {
		log.Fatal(err)
	}
	include, err := filepath.Glob(includePath + "*.tmpl")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range include {
		fileName := filepath.Base(file)
		files := append(layout, file)
		res, err := template.New(fileName).
			Funcs(template.FuncMap{"markdown": markdownTemplater}).
			ParseFiles(files...)
		if err != nil {
			log.Fatal(err)
		}
		res, err = res.Parse(tmplStr)
		if err != nil {
			log.Fatal(err)
		}

		allTemplates[fileName] = res
	}
}

func renderTemplate(w http.ResponseWriter, name string, data interface{}) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	err := allTemplates[name].ExecuteTemplate(w, "main", data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

}
