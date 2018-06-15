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
	return http.StripPrefix(path, fs)
}
