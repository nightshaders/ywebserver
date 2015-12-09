package server

import (
	"fmt"
	"path/filepath"
	"io/ioutil"
	"log"
	"github.com/russross/blackfriday"
	"github.com/nightshaders/ywebserver/config"
	"os"
)


func ProcessMarkdownTemplate(mdHtml []byte, conf *config.WebConf) ([]byte, error) {
	tpl, err := ioutil.ReadFile(conf.MarkdownTemplate)
	return tpl, err
}

func CompileMarkdown(path string) ([]byte, error) {
	dir, err := os.Getwd()
	if err != nil {
		log.Println(err)
		return nil, err
	}

	bb, err := ioutil.ReadFile(filepath.Join(dir, path))
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return blackfriday.MarkdownBasic(bb), nil
}

func HybridAssets(assets config.EmbeddedAsset, conf *config.WebConf) config.EmbeddedAsset {
	return func(path string) ([]byte, error) {
		ext := filepath.Ext(path)
		if IsMarkdown(ext) {
			md, err := CompileMarkdown(path)
			if err != nil {
				return nil, fmt.Errorf("Unable to compile markdown at %s", path)
			}
			return ProcessMarkdownTemplate(md, conf)
		} else {
			return assets(path)
		}
		return nil, fmt.Errorf("Have yet to implement %s", "HybridAssets")
	}
}

func IsMarkdown(ext string) bool {
	return ext == ".md"
}




