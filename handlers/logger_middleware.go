package handlers

import (
	"time"
	"github.com/gofiber/fiber/v2"
	"github.com/SimonBees/beeBot/utils"
)

// LoggerMiddleware logs all HTTP requests
func LoggerMiddleware(c *fiber.Ctx) error {
	start := time.Now()

	// Process request
	err := c.Next()

	// Calculate duration
	duration := time.Since(start)

	// Log the request
	utils.LogRequest(c.Method(), c.Path(), duration, c.Response().StatusCode())

	return err
}
