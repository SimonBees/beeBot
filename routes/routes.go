package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/SimonBees/beeBot/handlers"
)

func SetupRoutes(app *fiber.App) {
	// Base routes
	app.Get("/", handlers.WelcomeHandler)
	app.Get("/health", handlers.HealthHandler)
	
	// API routes group
	api := app.Group("/api/v1")
	
	// Authentication routes
	auth := api.Group("/auth")
	auth.Post("/register", handlers.RegisterHandler)
	auth.Post("/login", handlers.LoginHandler)
	
	// Protected routes
	protected := api.Group("/user").Use(handlers.ValidateJWTMiddleware)
	protected.Get("/profile", handlers.GetUserProfileHandler)
	
	// Add more routes here as needed
	setupUserRoutes(api)
}

func setupUserRoutes(api fiber.Router) {
	// User-related routes will be added here
}