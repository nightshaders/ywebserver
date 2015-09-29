package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/nightshaders/ywebserver/config"
	"github.com/nightshaders/ywebserver/embedded"
	"log"
	"net/http"
	"path/filepath"
	"strings"
)

func MimeType(ext string) (string, error) {
	switch ext {
	case "tff":
		return "application/x-font-ttf", nil
	case "js":
		return "application/javascript", nil
	case "json":
		return "application/json", nil
	case "css":
		return "text/css", nil
	case "html", "htm":
		return "text/html", nil
	default:
		return "", nil
	}
}

func EmbeddedAssetPath(wc *config.WebConf, assetPath string) string {
	if assetPath == "/" || assetPath == "" {
		assetPath = wc.DefaultFile
	}
	if strings.HasPrefix(assetPath, "/") && len(assetPath) > 1 {
		assetPath = assetPath[1:]
	}
	return assetPath
}

func EmbeddedCheck(wc *config.WebConf) func(r *http.Request, rm *mux.RouteMatch) bool {
	return func(r *http.Request, rm *mux.RouteMatch) bool {
		asset := EmbeddedAssetPath(wc, r.URL.Path)
		fmt.Printf("Finding resrouce: %s\n", asset)
		fileBytes, err := embedded.Asset(asset)
		exists := err == nil && fileBytes != nil && len(fileBytes) > 0
		return exists
	}
}

func ServeEmbedded(wc *config.WebConf) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		asset := EmbeddedAssetPath(wc, r.URL.Path)
		filebytes, err := embedded.Asset(asset)
		if err != nil {
			log.Printf("Didn't find embedded asset: %s\n", asset)
			http.NotFound(w, r)
			return
		}

		ext := filepath.Ext(r.URL.Path)
		mime, err := MimeType(ext)
		if err != nil {
			log.Printf("Didn't find extention: %s\n", ext)
			http.NotFound(w, r)
			return
		}

		w.Header().Set("Cache-Control", "public, max-age=315360000")
		w.Header().Set("Content-Type", mime)
		w.Write(filebytes)
	}
}
