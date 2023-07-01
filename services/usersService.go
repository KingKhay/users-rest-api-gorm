package services

import (
	"users-rest-api-gorm/dto"
	"users-rest-api-gorm/models"
	"users-rest-api-gorm/repository"
	"users-rest-api-gorm/utils"
)

type UserService struct {
	repository repository.UserRepository
}

func NewUserService() *UserService {
	return &UserService{repository: repository.UserRepository{}}
}

func (s *UserService) CreateUser(theUser *models.User) (*models.User, error) {
	hashedPassword, err := utils.HashPassword(theUser.Password)
	if err != nil {
		return nil, err
	}

	theUser.Password = hashedPassword
	createdUser, err := s.repository.CreateUser(*theUser)
	if err != nil {
		return nil, err
	}
	return createdUser, nil
}

func (s *UserService) GetAllUsers() ([]models.User, error) {
	users, err := s.repository.GetAllUsers()
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UserService) GetUserById(id int) (*models.User, error) {
	foundUser, err := s.repository.GetUserByID(id)
	if err != nil {
		return nil, err
	}
	return foundUser, nil
}

func (s *UserService) UpdateUser(id int, updated dto.UpdateUserDTO) (*models.User, error) {
	user, err := s.repository.UpdateUser(id, updated)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) DeleteUser(id int) error {
	err := s.repository.DeleteUser(id)
	if err != nil {
		return err
	}
	return nil
}
