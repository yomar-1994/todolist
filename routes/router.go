package routes

import (
	"todolist/handlers"

	"github.com/gorilla/mux"
)

func SetupRouter(UserHandler *handlers.UserHandler, todoHandler *handlers.TodoHandler) *mux.Router {
	r := mux.NewRouter()


    //  // Public routes 
	r.HandleFunc("/register", UserHandler.RegisterUser).Methods("POST")
	r.HandleFunc("/login", UserHandler.Login).Methods("POST")

	// //sub router for protected routes
	
	// protected := r.PathPrefix("/").Subrouter()
	// protected.Use(auth.AuthMiddleware)

	// //authenticated routes
	// protected.HandleFunc("/protect", handlers.ProtectedHandler).Methods("GET")
	// protected.HandleFunc("/todos", handlers.CreateTodo).Methods("POST")
	// protected.HandleFunc("/todos", handlers.GetTodos).Methods("GET")
	// protected.HandleFunc("/todos/{id}", handlers.UpdateTodo).Methods("PUT")
	// protected.HandleFunc("/todos/{id}", handlers.DeleteTodo).Methods("DELETE")

	return  r
}
