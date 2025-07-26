package repository

import (
	"todolist/db"
	"todolist/models"
)

type TodoRepository interface {
	GetTodoByName(name string) (*models.Todo, error)
	CreateTodo(todo *models.Todo) error
	GetAllTodos() ([]models.Todo,error)
}

type TodoRepo struct{}

func (r *TodoRepo) GetTodoByName(name string) (*models.Todo, error) {
	var todo models.Todo
	err := db.DB.Where("title = ?", name).First(&todo).Error
	if err != nil {
		return &models.Todo{}, err
	}

	return &todo, nil

}

func (r *TodoRepo) CreateTodo(todo *models.Todo) error {
	err := db.DB.Create(&todo).Error
	if err != nil {
		return  err
	}
	return  nil
	
}

func (r *TodoRepo) GetAllTodos() ([]models.Todo,error){
	var todos []models.Todo
	err := db.DB.Find(&todos).Error
	if err != nil {
		return nil, err
	}

	return todos, nil
}
