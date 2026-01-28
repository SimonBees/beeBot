package handlers

import (
	"github.com/gofiber/fiber/v2"
)

// HealthHandler returns the health status of the application
func HealthHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": "OK",
		"uptime": "running",
		"service": "beeBot Server",
	})
}

// WelcomeHandler returns a welcome message
func WelcomeHandler(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "Welcome to beeBot Server!",
		"version": "1.0.0",
		"status": "active",
	})
}