package services

import (
	"banking-system/internal/models"
	"banking-system/internal/repository"
	"fmt"
)

type UserService interface {
	ValidateUser(req *models.RegisterUserRequest) error
}
type userService struct {
	repo repository.UsersRepository
}

func NewUserService(r repository.UsersRepository) UserService {
	return &userService{repo: r}
}
func (s *userService) ValidateUser(req *models.RegisterUserRequest) error {
	valid, err := s.repo.GetUserByEmail(req.Email)
	if err != nil {
		return err
	}
	if !valid {
		return fmt.Errorf("user with email %s already exists", req.Email)
	} else {
		return s.repo.CreateUser()
	}
}
