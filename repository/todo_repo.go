package repository

import (
	"todolist/db"
	"todolist/models"
)

type TodoRepository interface {
	GetTodoByName(name string) (*models.Todolist, error)
	CreateTodo(todo *models.Todolist) error
}

type TodoRepo struct{}

func (r *TodoRepo) GetTodoByName(name string) (*models.Todolist, error) {
	var todo models.Todolist
	err := Db.DB.Where("name = ?", todo.Name).First(&todo).Error
	if err != nil {
		return &models.todolist{}, err
	}

	return &todo, nil

}

func (r *TodoRepo) CreateTodo(todo *models.Todolist) error {
	err := Db.DB.Create(&todo).Error
	if err != nil {
		return  err
	}
	return  nil
	
}
