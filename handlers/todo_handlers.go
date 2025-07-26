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
	
	var todo models.Todo
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

func (h *TodoHandler) GetTodos(w http.ResponseWriter, r *http.Request) {
	
    //call service layer directly 
	todos, err := h.Service.GetTodos()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//send reponse 
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(todos)
}