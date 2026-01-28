package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/SimonBees/beeBot/config"
	"github.com/SimonBees/beeBot/routes"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize Fiber app
	app := fiber.New(fiber.Config{
		AppName: "beeBot Server v1.0.0",
	})

	// Setup routes
	routes.SetupRoutes(app)

	// Start server
	port := ":" + cfg.Port
	log.Printf("beeBot Server starting on port %s", port)
	log.Fatal(app.Listen(port))
}