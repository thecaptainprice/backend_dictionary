// services/user_service.go

package services

import (
	"github.com/thecaptainprice/dictionary-app/backend/models"
	"github.com/thecaptainprice/dictionary-app/backend/repositories"
)

type UserService struct {
	UserRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{UserRepo: userRepo}
}

func (s *UserService) GetUserByID(id uint64) (*models.User, error) {
	return s.UserRepo.GetUserByID(id)
}

// GetUserByEmail retrieves a user from the database by email
func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	return s.UserRepo.GetUserByEmail(email)
}

// CreateUser creates a new user in the database
func (s *UserService) CreateUser(user *models.User) error {
	return s.UserRepo.CreateUser(user)
}

// UpdateUser updates an existing user in the database
func (s *UserService) UpdateUser(user *models.User) error {
	return s.UserRepo.UpdateUser(user)
}

// DeleteUser deletes an existing user from the database
func (s *UserService) DeleteUser(id uint64) error {
	return s.UserRepo.DeleteUser(id)
}
