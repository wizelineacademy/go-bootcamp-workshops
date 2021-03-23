package service

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"clean/model"
)

// Service
type Service struct {
	csvr *os.File
	csvw *csv.Writer
}

// New creates a new Service layer
func New(
	csvr *os.File,
	csvw *csv.Writer) (*Service, error) {

	return &Service{csvr, csvw}, nil
}

// GetTodos -
func (s *Service) GetTodos() ([]*model.Todo, error) {
	out := []*model.Todo{}

	csvr := csv.NewReader(s.csvr)

	data, err := csvr.ReadAll()
	if err != nil {
		fmt.Print("Error reading")
		return nil, err
	}

	for _, row := range data {
		todo := model.Todo{}

		id, _ := strconv.Atoi(row[0])
		todo.ID = id
		todo.Task = row[1]
		todo.Status = row[2]
		if row[3] == "false" {
			todo.IsDeleted = false
		} else {
			todo.IsDeleted = true
		}

		out = append(out, &todo)
	}
	s.csvr.Seek(0, 0)

	return out, nil
}

// GetTodo -
func (s *Service) GetTodo(todoID string) (*model.Todo, error) {
	out := &model.Todo{}

	csvr := csv.NewReader(s.csvr)

	data, err := csvr.ReadAll()
	if err != nil {
		fmt.Print("Error reading")
		return nil, err
	}

	for _, row := range data {
		if row[0] == todoID {
			id, _ := strconv.Atoi(row[0])
			out.ID = id
			out.Task = row[1]
			out.Status = row[2]
			if row[3] == "false" {
				out.IsDeleted = false
			} else {
				out.IsDeleted = true
			}
		}
	}
	s.csvr.Seek(0, 0)

	return out, nil
}
