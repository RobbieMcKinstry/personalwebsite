package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/russross/blackfriday"
)

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
