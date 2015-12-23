package server

import (
	"net/http"
	"path"

	"github.com/gorilla/mux"
)

func MatchAssets(req *http.Request, rm *mux.RouteMatch) bool {
	p := req.URL.Path
	e := path.Ext(p)

	// Typically handled by index.html, but that is set in the Conf
	if p == "/" {
		return true
	}

	switch e {
	case ".png", ".jpg", ".jpeg", ".gif", ".ico":
		return true
	case ".js", ".css", ".html", ".ttf", ".json":
		return true
	default:
		return false
	}
}
