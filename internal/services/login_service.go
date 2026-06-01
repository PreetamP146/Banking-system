package services

import (
	"banking-system/internal/models"
	"banking-system/pkg/utils"
	"fmt"
)

func (s *userService) LoginUser(req *models.LoginUserRequest) (string, error) {
	user, err := s.repo.GetUserByEmail(req.Email)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", fmt.Errorf("Invalid email or password")
	}
	if !utils.CheckPasswordHash(req.Password, user.Password_Hash) {
		return "", fmt.Errorf("Invalid email or password")
	}
	token, err := utils.GenerateJWT(user.ID, user.Role, s.cfg)
	if err != nil {
		return "", err
	}
	return token, nil
}
