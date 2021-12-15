package usecases

import "github.com/pinghsuanC/todoapp_react_and_golang/backend/entities"

func GetTodos(repo TodosRepository) ([]entities.Todo, error) {
	return nil, ErrInternal;
}