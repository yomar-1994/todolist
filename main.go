package main

import (
	"fmt"
	"log"
	"net/http"
	"todolist/db"
	"todolist/handlers"
	"todolist/repository"
	"todolist/routes"
	"todolist/services"
)

func main() {
	db.InitDb()

	//initialize repositories
	UserRepo := &repository.UserRepo{}
	todoRepo := &repository.TodoRepo{}

	// initialize service
	userService := &services.UserService{Repo: UserRepo}
	TodoService := &services.TodoServioce{Repo: todoRepo}

	//initialize handlers
	UserHandler := &handlers.UserHandler{Service: userService}
	todoHandler := &handlers.TodoHandler{Service: TodoService}

	//routes
	routes.SetupRouter(UserHandler, todoHandler)

	//start server
	fmt.Println("server started on :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("failed to start server", err)
	}

}
