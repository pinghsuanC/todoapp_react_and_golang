package usecases_test

import (
	"fmt"
	"testing"

	"github.com/pinghsuanC/todoapp_react_and_golang/backend/entities"
	"github.com/pinghsuanC/todoapp_react_and_golang/backend/usecases"
)
type MockTodosRepo struct {}
func (MockTodosRepo) GetAllTodos() ([]entities.Todo, error){
	return nil, fmt.Errorf("Something went wrong!")
};

func TestGetTodos(t *testing.T) {
	// think about error cases
	repo := new(MockTodosRepo)

	todos, err := usecases.GetTodos(repo)

	if err != usecases.ErrInternal {
		t.Fatalf("Expected ErrorInternal; Got: %v", err)
	}

	if todos !=nil {
		t.Fatalf("Expected todos to be nil, got: %v", todos)
	}

}