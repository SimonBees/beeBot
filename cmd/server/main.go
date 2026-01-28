package main

import (
	"log"
	"github.com/gofiber/fiber/v2"
	"github.com/SimonBees/beeBot/config"
	"github.com/SimonBees/beeBot/routes"
	"github.com/SimonBees/beeBot/handlers"
	"github.com/SimonBees/beeBot/utils"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize logger
	utils.InitLogger()
	utils.LogInfo("Starting beeBot Server...")

	// Initialize database
	config.InitDatabase()

	// Initialize Fiber app
	app := fiber.New(fiber.Config{
		AppName: "beeBot Server v1.0.0",
	})

	// Apply middlewares
	app.Use(handlers.LoggerMiddleware)

	// Setup routes
	routes.SetupRoutes(app)

	// Start server
	port := ":" + cfg.Port
	log.Printf("beeBot Server starting on port %s", port)
	utils.LogInfo("Server listening on port " + cfg.Port)
	log.Fatal(app.Listen(port))
}