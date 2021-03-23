package usecase

import (
	"clean/model"
)

// UseCase struct
type UseCase struct {
	service Service
}

// Service interface
type Service interface {
	GetTodos() ([]*model.Todo, error)
	GetTodo(todoID string) (*model.Todo, error)
}

// New UseCase
func New(service Service) *UseCase {
	return &UseCase{service}
}

// GetTodos -
func (u *UseCase) GetTodos() ([]*model.Todo, error) {
	resp, err := u.service.GetTodos()

	if err != nil {
		return nil, err
	}

	return resp, nil
}

// GetTodo -
func (u *UseCase) GetTodo(todoID string) (*model.Todo, error) {
	resp, err := u.service.GetTodo(todoID)

	if err != nil {
		return nil, err
	}

	return resp, nil
}
