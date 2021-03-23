package controller

import (
	"net/http"

	"clean/model"

	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// UseCase interface
type UseCase interface {
	GetTodos() ([]*model.Todo, error)
	GetTodo(string) (*model.Todo, error)
}

// TodoUsecase struct
type TodoUsecase struct {
	useCase UseCase
	render  *render.Render
}

// New returns a controller
func New(
	u UseCase,
	r *render.Render,
) *TodoUsecase {
	return &TodoUsecase{u, r}
}

// GET /todos
func (t *TodoUsecase) GetTodos(w http.ResponseWriter, r *http.Request) {
	body, _ := t.useCase.GetTodos()
	w.Header().Set("Content-Type", "application/json")
	t.render.JSON(w, http.StatusOK, body)
}

// GET /todos/{id}
func (t *TodoUsecase) GetTodo(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	body, _ := t.useCase.GetTodo(params["id"])
	w.Header().Set("Content-Type", "application/json")
	t.render.JSON(w, http.StatusOK, body)
}
