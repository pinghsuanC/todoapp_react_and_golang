package usecases

import "github.com/pinghsuanC/todoapp_react_and_golang/backend/entities"

type TodosRepository interface {
	GetAllTodos() ([]entities.Todo, error)
}