package handler

import (
	"banking-system/internal/models"
	"banking-system/internal/services"
	"banking-system/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	RegisterUser(c *fiber.Ctx) error
}
type userHandler struct {
	svc services.UserService
}

func NewUserHandler(s services.UserService) UserHandler {
	return &userHandler{svc: s}
}
func (h *userHandler) RegisterUser(c *fiber.Ctx) error {
	var req models.RegisterUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
	if !utils.ValidateEmail(req.Email) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid email format"})
	}
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to hash password"})
	}
	req.Password = hashedPassword
	if err := h.svc.ValidateUser(&req); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to validate user"})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User registered successfully"})
}
