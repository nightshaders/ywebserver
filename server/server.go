package server

import (
	"github.com/gorilla/mux"
	"github.com/nightshaders/ywebserver/config"
	"net/http"
)

type Server struct {
	*mux.Router
	Conf *config.WebConf
}

func NewServer(conf *config.WebConf) *Server {
	return &Server{
		Router: mux.NewRouter(),
		Conf:   conf,
	}
}

func (s *Server) HandleHttp(path string, h http.Handler) *Server {
	s.HandleFunc(path, h.ServeHTTP)
	return s
}
