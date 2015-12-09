package config

import (
	"encoding/json"
	"fmt"
	"log"
)


type WebConf struct {
	Port                 int
	SiteRoot             string
	DefaultFile          string
	AssetPath            string
	ServeEmbedddedAssets bool
	EmbeddedAsset        EmbeddedAsset      `json:"-"`
	MarkdownTemplate     string
}

type EmbeddedAsset func(name string) ([]byte, error)

func (w *WebConf) String() string {
	res, err := json.MarshalIndent(w, "", "   ")
	if err != nil {
		log.Println(err)
	}
	return string(res)
}

func defInt(a, b int) int {
	if a == 0 {
		return b
	} else {
		return a
	}
}

func defString(a, b string) string {
	if a == "" {
		return b
	} else {
		return a
	}
}

func (w *WebConf) Host() string {
	return fmt.Sprintf(":%d", w.Port)
}

func (w *WebConf) ApplyDefaults() *WebConf {
	def := defaultWebConf
	rs := &WebConf{
		Port:                 defInt(w.Port, def.Port),
		SiteRoot:             defString(w.SiteRoot, def.SiteRoot),
		DefaultFile:          defString(w.DefaultFile, def.DefaultFile),
		ServeEmbedddedAssets: w.ServeEmbedddedAssets,
	}
	return rs
}

var defaultWebConf *WebConf = &WebConf{
	Port:        9100,
	SiteRoot:    "src/web_pair/.www/dest/_site",
	DefaultFile: "index.html",
}
