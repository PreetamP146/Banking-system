package routes

import (
	"banking-system/internal/handler"
	"banking-system/pkg/config"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(api fiber.Router, handler handler.UserHandler, cfg *config.Config) {
	auth := api.Group("/auth")
	auth.Post("/register", handler.RegisterUser)
}
