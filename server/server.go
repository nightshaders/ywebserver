package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nightshaders/ywebserver/config"
	"github.com/nightshaders/ywebserver/decorator"
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

func (r *Server) Route(route string, pipe decorator.Decorator, h decorator.Handler) {
	r.HandleFunc(route, pipe.Handle(h))
}

func (r *Server) DefaultServeStatic() *mux.Route {
	if r.Conf.ServeEmbedddedAssets {
		return r.MatcherFunc(EmbeddedCheck(r.Conf)).HandlerFunc(ServeEmbedded(r.Conf))
	} else {
		return r.MatcherFunc(MatchAssets).HandlerFunc(ServeFile(r.Conf))
	}
}
