package service

import (
	"a2_ems_project/model"
	"a2_ems_project/repository"
	"fmt"
)

type UserService struct {
	UserRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{UserRepo: userRepo}
}

// Register a new user
func (service *UserService) RegisterUser(user *model.User) error {
	// Check if the username already exists
	exists, err := service.UserRepo.IsUsernameTaken(user.Username)
	if err != nil {
		return fmt.Errorf("error checking username: %v", err)
	}
	if exists {
		return fmt.Errorf("username '%s' is already taken", user.Username)
	}

	// Create the user
	err = service.UserRepo.CreateUser(user)
	if err != nil {
		return fmt.Errorf("could not register user: %v", err)
	}
	return nil
}

// Authenticate a user
func (service *UserService) AuthenticateUser(username, password string) (*model.User, error) {
	return service.UserRepo.AuthenticateUser(username, password)
}
