package main

import (
	"banking-system/internal/configdb"
	"banking-system/internal/routes"
	"banking-system/pkg/config"
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Load configuration
	cfg := config.LoadConfig()

	// Database Connection
	db := configdb.LoadDBConfig(cfg)

	// test database connection
	if err := db.PingContext(ctx); err != nil {
		log.Fatalf("Database connection failed: %v", err)
	}
	log.Println("Database connection established")
	//create fiber app
	app := fiber.New()
	// Setup routes
	routes.SetupRoutes(app, db, cfg)
	// Create a simple health check endpoint
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.SendString("Server is healthy!")
	})

	// Start the server
	log.Println("Server running on port :", cfg.AppPort)
	if err := app.Listen(":" + cfg.AppPort); err != nil {
		log.Fatalf("Error in server starting: %v", err)
	}
}
