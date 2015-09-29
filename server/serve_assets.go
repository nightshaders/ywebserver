package server

import (
	"fmt"
	"github.com/nightshaders/ywebserver/config"
	"log"
	"net/http"
	"os"
	"path/filepath"
)

func FileExists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil || !os.IsNotExist(err)
}

func FileToServe(wc *config.WebConf, r *http.Request) string {
	path := r.URL.Path
	log.Printf("serving path: %s %v", path, path == "/")

	if r.URL.Path == "/" || r.URL.Path == "" {
		path = wc.DefaultFile
		log.Printf("serving default file: %s", path)
	}

	file := filepath.Join(wc.SiteRoot, path)
	return file
}

func ServeFile(wc *config.WebConf) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		filename, err := filepath.Abs(FileToServe(wc, r))
		if err != nil {
			fmt.Println(err)
		}
		exists := FileExists(filename)

		log.Printf("filename: %s, exists: %v", filename, exists)

		if exists {
			http.ServeFile(w, r, filename)
		} else {
			http.NotFound(w, r)
		}
	}
}
