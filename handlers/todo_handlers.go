package handlers

import (
	"encoding/json"
	"net/http"
	"todolist/models"
	"todolist/services"
)

type TodoHandler struct {
	Service *services.TodoService
}

//POST /todos: create a todo (authenticated).
func (h *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	
	var todo models.Todolist
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// call service layer 
	err = h.Service.CreateTodo(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}


	//response 
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}