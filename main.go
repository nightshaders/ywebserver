package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/nightshaders/ywebserver/decorator"
)

func main() {
	//	server.NewCli(ExampleServe).Run(os.Args)

	var unauth decorator.Decorator = decorator.NewPipeline().
		Next(ProvideErrorHandling).
		Next(ProvideSession).
		Next(ProvideUserProfile)

	handler := unauth.Handle(H{})

	res := httptest.NewRecorder()
	req := &http.Request{}

	handler(res, req)
}

type H struct{}

func (h H) Handle(p decorator.Params) error {
	fmt.Println("Handle")
	return nil
}

func ProvideErrorHandling(w decorator.WebHandler) decorator.WebHandler {
	return func(p decorator.Params) (err error) {
		fmt.Println("ErrorHandling")
		return w.Handle(p)
	}
}

func ProvideSession(w decorator.WebHandler) decorator.WebHandler {
	return func(p decorator.Params) (err error) {
		fmt.Println("Session")
		return w.Handle(p)
	}
}

func ProvideUserProfile(w decorator.WebHandler) decorator.WebHandler {
	return func(p decorator.Params) (err error) {
		fmt.Println("UserProfile")
		return w.Handle(p)
	}
}
