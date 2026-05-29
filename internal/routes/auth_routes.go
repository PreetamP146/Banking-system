package routes

import (
	"banking-system/internal/handler"

	"github.com/gofiber/fiber/v2"
)

func AuthRoutes(api fiber.Router, handler handler.UserHandler) {
	auth := api.Group("/auth")
	auth.Post("/register", handler.RegisterUser)
	auth.Post("/login", handler.LoginUser)
}
