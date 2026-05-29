package handler

import (
	"banking-system/internal/models"
	"banking-system/pkg/utils"

	"github.com/gofiber/fiber/v2"
)

func (h *userHandler) LoginUser(c *fiber.Ctx) error {
	var req models.LoginUserRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request body"})
	}
	if !utils.ValidateEmail(req.Email) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid email format"})
	}
	token, err := h.svc.LoginUser(&req)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"token": token})
}
