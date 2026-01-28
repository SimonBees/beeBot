package handlers

import (
	"strings"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/SimonBees/beeBot/config"
	"github.com/SimonBees/beeBot/models"
)

// ValidateJWTMiddleware validates the JWT token and extracts user info
func ValidateJWTMiddleware(c *fiber.Ctx) error {
	authHeader := c.Get("Authorization")
	if authHeader == "" {
		return c.Status(401).JSON(fiber.Map{
			"error": "Authorization header missing",
		})
	}

	// Extract token from "Bearer <token>" format
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")
	if tokenString == authHeader {
		return c.Status(401).JSON(fiber.Map{
			"error": "Invalid authorization format. Use: Bearer <token>",
		})
	}

	// Parse and validate token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.ErrUnauthorized
		}
		return []byte(config.LoadConfig().JWTSecret), nil
	})

	if err != nil || !token.Valid {
		return c.Status(401).JSON(fiber.Map{
			"error": "Invalid or expired token",
		})
	}

	// Extract claims
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return c.Status(401).JSON(fiber.Map{
			"error": "Invalid token claims",
		})
	}

	// Get user ID from claims
	userID, ok := claims["user_id"].(float64)
	if !ok {
		return c.Status(401).JSON(fiber.Map{
			"error": "Invalid user ID in token",
		})
	}

	// Fetch user from database
	db := config.GetDB()
	var user models.User
	if err := db.First(&user, uint(userID)).Error; err != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": "User not found",
		})
	}

	// Store user in context
	c.Locals("user", &user)

	return c.Next()
}
