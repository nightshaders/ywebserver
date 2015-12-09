package main

import (
	"github.com/nightshaders/ywebserver/handlers"
	"github.com/nightshaders/ywebserver/server"
	"fmt"
	"log"
	"net/http"
)


// ExampleServe is an example function.  A similar function should
// be used to write a specific server.
func ExampleServe(r *server.Server) {
	fmt.Println(r.Conf.String())
	fmt.Println("Starting Web Server")
	r.Conf.EmbeddedAsset = server.HybridAssets(Asset, r.Conf)

	r.HandleFunc("/user", handlers.CreateUser).Methods("POST")
	r.HandleFunc("/users", handlers.GetUsers).Methods("GET")
	r.HandleFunc("/hello/{name}", handlers.Hello).Methods("GET")
	r.HandleHttp("/x", &handlers.HandlerX{"Superman"})

	// Handling static assets
	if r.Conf.ServeEmbedddedAssets {
		r.MatcherFunc(server.EmbeddedCheck(r.Conf)).
		HandlerFunc(server.ServeEmbedded(r.Conf))
	} else {
		r.MatcherFunc(server.MatchAssets).
		HandlerFunc(server.ServeFile(r.Conf))
	}

	err := http.ListenAndServe(r.Conf.Host(), r)
	if err != nil {
		log.Fatal(err)
	}
}

