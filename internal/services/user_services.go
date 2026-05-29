package services

import (
	"banking-system/internal/models"
	"banking-system/internal/repository"
	"banking-system/pkg/utils"
	"fmt"
)

type UserService interface {
	RegisterUser(req *models.RegisterUserRequest) error
}
type userService struct {
	repo repository.UsersRepository
}

func NewUserService(r repository.UsersRepository) UserService {
	return &userService{repo: r}
}
func (s *userService) RegisterUser(req *models.RegisterUserRequest) error {
	valid, err := s.repo.IsUniqueEmail(req.Email)
	if err != nil {
		return err
	}
	if !valid {
		return fmt.Errorf("User with email %s already exists", req.Email)
	} else {
		hash, err := utils.HashPassword(req.Password)
		if err != nil {
			return err
		}

		req.Password = hash
		return s.repo.CreateUser(req)
	}
}
