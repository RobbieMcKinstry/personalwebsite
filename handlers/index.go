package handlers

import (
	"net/http"
)

// The IndexHandler points to the root of the application.
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	renderTemplate(w, "index.tmpl", nil)
}
