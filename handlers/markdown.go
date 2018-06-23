package handlers

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"path/filepath"

	"github.com/russross/blackfriday"
)

const markdownRoot = "markdown"

func markdownTemplater(args ...interface{}) template.HTML {
	s := blackfriday.MarkdownCommon([]byte(fmt.Sprintf("%s", args...)))
	return template.HTML(s)
}

func renderMarkdown(w http.ResponseWriter, r *http.Request, title, body string) {
	renderTemplate(w, "markdown.tmpl", markdownPost{title, body})
}

type markdownPost struct {
	Title, Body string
}

// RenderMarkdown looks into the markdown directory for a .md with the provided filename
// and renders it.
func RenderMarkdown(w http.ResponseWriter, r *http.Request, title, filename string) {
	var content, err = loadMarkdownFromFilesystem(filename)
	if err != nil {
		panic("No such markdown file")
	}
	renderMarkdown(w, r, title, content)
}

func loadMarkdownFromFilesystem(path string) (string, error) {
	var filename = filepath.Join(markdownRoot, path)
	var contents, err = ioutil.ReadFile(filename)
	return string(contents), err
}
