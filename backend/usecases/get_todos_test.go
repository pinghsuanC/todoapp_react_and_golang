package usecases_test

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/pinghsuanC/todoapp_react_and_golang/backend/entities"
	"github.com/pinghsuanC/todoapp_react_and_golang/backend/usecases"
)

var dummyTodos = []entities.Todo{
	{Title: "todo 1", Description: "descriotion 1", IsCompleted: false},
	{Title: "todo 2", Description: "descriotion 2", IsCompleted: true},
	{Title: "todo 3", Description: "descriotion 3", IsCompleted: false},
};

type BadTodosRepo struct {}
func (BadTodosRepo) GetAllTodos() ([]entities.Todo, error){
	return nil, fmt.Errorf("Something went wrong!")
};
type MockTodosRepo struct {}
func (MockTodosRepo) GetAllTodos() ([]entities.Todo, error){
	return dummyTodos, nil
};

func TestGetTodos(t *testing.T) {
	// case 1
	t.Run("Returns ErrInternal when TodosRepository returns error", func(t *testing.T){
		// think about error cases
		repo := new(BadTodosRepo)
		todos, err := usecases.GetTodos(repo)

		if err != usecases.ErrInternal {
			t.Fatalf("Expected ErrorInternal; Got: %v", err)
		}

		if todos !=nil {
			t.Fatalf("Expected todos to be nil, got: %v", todos)
		}
	})

	// test 2
	t.Run("Returns todos from TodosRepository", func(t *testing.T){
		repo := new(MockTodosRepo)
		todos, err := usecases.GetTodos(repo)

		if err != nil {
			t.Fatalf("Expected nil; Got: %v", err)
		}

		if !reflect.DeepEqual(todos, dummyTodos) {
			t.Fatalf("Expected todos to be an array, got: %v", todos)
		}
	})
	
}