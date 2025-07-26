package services

import (
	"todolist/models"
	"todolist/repository"

	"github.com/google/uuid"
)

type TodoService struct {
	Repo repository.TodoRepository
}

func (s *TodoService) CreateTodo(req *models.Todo) error {

	// check it doesnt exist in the db already
	_, err := s.Repo.GetTodoByName(req.Title)
	if err == nil {
		return err
	}

	myuuid := uuid.NewString()
	req.ID = myuuid

	// add to db using encoder
	err = s.Repo.CreateTodo(req)
	if err != nil {
		return err
	}
	return nil

}
