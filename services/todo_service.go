package services

import (
	"todolist/models"
	"todolist/repository"

	"github.com/google/uuid"
)

type TodoServioce struct {
	Repo repository.TodoRepository
}

func (s *TodoServioce) CreateTodo(req models.Todolist) error {

	// check it doesnt exist in the db already
	_, err := s.Repo.GetTodoByName(req.Name)
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
