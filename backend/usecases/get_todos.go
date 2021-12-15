package usecases

import "github.com/pinghsuanC/todoapp_react_and_golang/backend/entities"

func GetTodos(repo TodosRepository) ([]entities.Todo, error) {
	todos, err := repo.GetAllTodos()
	if(err!=nil){
		return nil, ErrInternal;
	}
	return todos, nil;
}