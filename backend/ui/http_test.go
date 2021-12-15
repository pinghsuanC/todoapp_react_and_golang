package ui_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/pinghsuanC/todoapp_react_and_golang/backend/entities"
	"github.com/pinghsuanC/todoapp_react_and_golang/backend/ui"
)


func TestHTTP(t *testing.T){
	tests := getTest();
	tests = append(tests, getDisallowedMthodTests()...);
	testHTTP(tests, t)
	
}

type MockService struct {
	err   error
	todos []entities.Todo
}

type HTTPTest struct{
	name string
	service *MockService
	inputMethod string
	inputURL string

	expectedStatus int
	expectedTodos []entities.Todo
}

var dummyTodos = []entities.Todo{
	{Title: "todo 1", Description: "descriotion 1", IsCompleted: false},
	{Title: "todo 2", Description: "descriotion 2", IsCompleted: true},
	{Title: "todo 3", Description: "descriotion 3", IsCompleted: false},
};

func (s MockService) GetAllTodos()([]entities.Todo, error){
	if s.err != nil{
		return nil, s.err
	}
	return s.todos, nil
}



func testHTTP(tests []HTTPTest, t *testing.T)  {
	for _, test := range tests{
		t.Run(test.name, func(t *testing.T){
			service := &MockService{err: fmt.Errorf("something bad happened")}

			w := httptest.NewRecorder()
			r := httptest.NewRequest(test.inputMethod, test.inputURL, nil)

			server := ui.NewHTTP();
			server.UseService(service)
			server.ServeHTTP(w, r)
			
			var body []entities.Todo
			json.NewDecoder(w.Result().Body).Decode(&body)

			if(w.Result().StatusCode!=test.expectedStatus){
				t.Fatalf("Expected status to be %v, got: %v", test.expectedStatus, w.Result().StatusCode)
			}
			if(!reflect.DeepEqual(body, test.expectedTodos)){
				t.Fatalf("Expected todos to be an %v, got: %v", test.expectedTodos, body)
			}
		})
	}
}

func getTest() []HTTPTest {
	return []HTTPTest{
		{
			name: "Random error gives 500 status and no todos",
			service: &MockService{err:fmt.Errorf("something bad happened")}, 
			inputMethod: "GET",
			inputURL:"http://mywebsite.com/todos",
			expectedStatus: 500,
			expectedTodos: nil,
		},
		{
			name: "Random error gives 500 status and no todos",
			service: &MockService{err:fmt.Errorf("something bad happened")}, 
			inputMethod: "GET",
			inputURL:"http://mywebsite.com/todos/",
			expectedStatus: 500,
			expectedTodos: nil,
		},
		{
			name: "Wrong path gives 404 status and no todos",
			service: &MockService{todos: dummyTodos}, 
			inputMethod: "GET",
			inputURL:"http://mywebsite.com/foo/",
			expectedStatus: 404,
			expectedTodos: nil,
		},
		{
			name: "Wrong path gives 404 status and no todos",
			service: &MockService{todos: dummyTodos}, 
			inputMethod: "GET",
			inputURL:"http://mywebsite.com/test/",
			expectedStatus: 404,
			expectedTodos: nil,
		},
		{
			name: "Wrong path gives 404 status and no todos",
			service: &MockService{todos: dummyTodos}, 
			inputMethod: "GET",
			inputURL:"http://mywebsite.com/aaa",
			expectedStatus: 404,
			expectedTodos: nil,
		},
		{
			name: "Wrong method gives 405 status and no todos",
			service: &MockService{err:fmt.Errorf("405")}, 
			inputMethod: "GET",
			inputURL:"http://mywebsite.com/aaa",
			expectedStatus: 404,
			expectedTodos: nil,
		},
	}
}

func getDisallowedMthodTests() []HTTPTest {
	
	disAllowedMthods := []string{
		http.MethodDelete,
		http.MethodHead,
		http.MethodOptions,
		http.MethodPatch,
		http.MethodPost,
		http.MethodPut,
	}
	
	disallowedTest := []HTTPTest{};
	for _, method := range disAllowedMthods {
		
		disallowedTest = append(disallowedTest, HTTPTest{
			name: fmt.Sprintf("Method %s gives 405 status and no todos", method),
			service: &MockService{todos: dummyTodos},
			inputURL: "http://mywebsite.com/todos/",
			inputMethod: method,
			expectedStatus: http.StatusMethodNotAllowed,

		})

	}

	return disallowedTest;
}