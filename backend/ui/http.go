package ui

import (
	"encoding/json"
	"net/http"
)

type HTTPServer struct{
	service Service
}

func NewHTTP() *HTTPServer {
	return &HTTPServer{}
}

func (server *HTTPServer) UseService(s Service) {
	server.service = s;
}

func (server HTTPServer) ServeHTTP(w http.ResponseWriter, r *http.Request){
	verified := verifyRequest(w, r)

	if(!verified){
		return
	}

	todos, err := server.service.GetAllTodos()
	if(err!=nil && todos==nil){
		w.WriteHeader(http.StatusInternalServerError)
		return;
	}
	
	json.NewEncoder(w).Encode(todos)
}

func verifyRequest(w http.ResponseWriter, r *http.Request) bool {
	if(r.URL.Path != "/todos/" && r.URL.Path != "/todos"){
		w.WriteHeader(http.StatusNotFound)
		return false;
	}
	if(r.Method != http.MethodGet){
		w.WriteHeader(http.StatusMethodNotAllowed)
		return false;
	}
	return true;
}