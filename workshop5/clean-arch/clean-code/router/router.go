package router

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Controller interface {
	GetTodos(w http.ResponseWriter, r *http.Request)
	GetTodo(w http.ResponseWriter, r *http.Request)
}

func New(c Controller) *mux.Router {
	r := mux.NewRouter()

	// Endpoints
	r.HandleFunc("/todos", c.GetTodos).Methods("GET")
	r.HandleFunc("/todos/{id}", c.GetTodo).Methods("GET")

	return r
}
