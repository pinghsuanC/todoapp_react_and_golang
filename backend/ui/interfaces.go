package ui

import "github.com/pinghsuanC/todoapp_react_and_golang/backend/entities"

type Service interface {
	GetAllTodos() ([]entities.Todo, error)
}
