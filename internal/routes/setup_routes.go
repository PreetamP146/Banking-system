package routes

import (
	"banking-system/internal/config"
	"banking-system/internal/handler"
	"banking-system/internal/repository"
	"banking-system/internal/services"
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, db *sql.DB, cfg *config.Config) {
	userRepo := repository.NewUsersRepository(db)
	userService := services.NewUserService(userRepo, cfg)
	userHandler := handler.NewUserHandler(userService)

	api := app.Group("/api/v1/")
	AuthRoutes(api, userHandler)
}
