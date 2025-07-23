package repository

import (
	"todolist/models"
)
type UserRepository interface {
	GetUserByUsername(username string) (*models.User, error)
	CreateUser(user *models.User) error
}

type UserRepo struct{}

// handlers layer ----->services layer ----->repository layer (db methods)

func (r *UserRepo) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	err := db.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		return &models.User{}, err

	}

	return  &user, nil 
}

func (r *UserRepo) CreateUser(user *models.User) error {
	
	err := db.DB.Create(&user).Error
	if err != nil {
		return  err
	}

	return  nil 


}
