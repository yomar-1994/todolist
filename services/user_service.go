package services

import (
	"todolist/middleware"
	"todolist/models"
	"todolist/repository"
	"todolist/utils"

	"github.com/google/uuid"
)

type UserService struct {
	Repo repository.UserRepository
}

func (s *UserService) RegisterUser(req *models.User) error {

	// check if they exist in db
	_, err := s.Repo.GetUserByUsername(req.Username)
	if err == nil {
		return err
	}

	// hash the password
	HashPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return err

	}
	req.Password = HashPassword

	myuuid := uuid.NewString()
	req.ID = myuuid

	// put into db
	err = s.Repo.CreateUser(req)
	if err != nil {
		return err
	}

	return nil

}

func (s *UserService) Login(req *models.User) (string, error) {
	// check if user exists
	user, err := s.Repo.GetUserByUsername(req.Username)
	if err != nil {
		return "", err
	}

	//compare password
	err = utils.ComparePassword(user.Password, req.Password)
	if err != nil {
		return "", err
	}

	// generate token
	token, err := middleware.GenerateJWT(user.ID)
	if err != nil {
		return "", err
	}

	return token, nil
}
