package server

import (
	"github.com/gorilla/mux"
	"net/http"
	"ywebserver/config"
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
