package repository

import (
	"users-rest-api-gorm/dto"
	"users-rest-api-gorm/initializers"
	"users-rest-api-gorm/models"
)

type Repository interface {
	GetAllUsers() ([]models.User, error)
	GetUserByID(userID int) (*models.User, error)
	CreateUser(user models.User) (*models.User, error)
	UpdateUser(id int, user dto.UpdateUserDTO) (*models.User, error)
	DeleteUser(userID int) error
}

type UserRepository struct{}

func (s *UserRepository) GetAllUsers() ([]models.User, error) {
	var users []models.User
	err := initializers.DB.Find(&users).Error

	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserRepository) GetUserByID(id int) (*models.User, error) {
	var user models.User
	err := initializers.DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UserRepository) CreateUser(user models.User) (*models.User, error) {
	err := initializers.DB.Create(&user).Error

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *UserRepository) UpdateUser(id int, updatedUser dto.UpdateUserDTO) (*models.User, error) {
	var user models.User
	err := initializers.DB.First(&user, id).Error
	if err != nil {
		return nil, err
	}

	user.Name = updatedUser.Name
	user.Age = updatedUser.Age
	user.Email = updatedUser.Email

	err = initializers.DB.Save(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (s *UserRepository) DeleteUser(id int) error {
	return initializers.DB.Delete(&models.User{}, id).Error
}
