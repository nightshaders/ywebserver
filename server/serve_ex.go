package server

import (
	"fmt"
	"log"
	"net/http"
	"ywebserver/handlers"
)

func (r *Server) Serve() {
	conf := r.conf.ApplyDefaults()

	fmt.Println(r.conf.String())
	fmt.Println("Web Server Started")

	r.HandleFunc("/user", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	r.HandleFunc("/hello/{name}", handlers.Hello).Methods("GET")
	r.HandleHttp("/x", &handlers.HandlerX{"Superman"})

	// Handle static assets
	//r.MatcherFunc(MatchAssets).HandlerFunc(ServeFile(conf))
	r.MatcherFunc(EmbeddedCheck(conf)).HandlerFunc(ServeEmbedded(conf))

	err := http.ListenAndServe(conf.Host(), r)
	if err != nil {
		log.Fatal(err)
	}
}
