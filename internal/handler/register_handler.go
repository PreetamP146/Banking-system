package handler

import (
	"banking-system/internal/models"
	"banking-system/internal/services"
	"banking-system/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

type UserHandler interface {
	RegisterUser(c *fiber.Ctx) error
	LoginUser(c *fiber.Ctx) error
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
	if err := h.svc.RegisterUser(&req); err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User registered successfully"})
}
