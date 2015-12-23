package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/nightshaders/ywebserver/models"
)

type HandlerX struct {
	Name string
}

func (h *HandlerX) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "From HandlerX: %s", h.Name)
}

func toJson(u interface{}) string {
	b, _ := json.Marshal(u)
	s := string(b)
	return s
}

func name() string {
	return fmt.Sprintf("%s", "World")
}

func Hello(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	name := vars["name"]
	fmt.Fprintf(w, "Hello, %s!", name)
}

func GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, toJson(models.Users))
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err == nil {
		user := models.User{}
		json.Unmarshal(body, &user)
		models.Users = append(models.Users, user)
	} else {
		fmt.Println(err)
	}
}
