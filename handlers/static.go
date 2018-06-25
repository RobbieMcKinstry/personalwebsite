package handlers

import (
	"fmt"
	"net/http"
)

const staticPrefix = "static"

// Static returns the handle for all of the static assets
// used by this website.
func Static() http.Handler {
	return fsHandler(staticPrefix)
}

func fsHandler(dirName string) http.Handler {
	fs := http.FileServer(http.Dir(dirName))
	path := fmt.Sprintf("/%s/", dirName)
	var handler http.HandlerFunc = func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "max-age=3600")
		http.StripPrefix(path, fs).ServeHTTP(w, r)
	}
	return handler
}
