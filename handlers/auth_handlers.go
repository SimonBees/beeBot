package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/SimonBees/beeBot/models"
	"github.com/SimonBees/beeBot/config"
	"golang.org/x/crypto/bcrypt"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// RegisterHandler handles user registration
func RegisterHandler(c *fiber.Ctx) error {
	db := config.GetDB()
	
	var req models.RegisterRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Check if user already exists
	var existingUser models.User
	if err := db.Where("username = ? OR email = ?", req.Username, req.Email).First(&existingUser).Error; err == nil {
		return c.Status(409).JSON(fiber.Map{
			"error": "User with this username or email already exists",
		})
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Could not hash password",
		})
	}

	// Create user
	user := models.User{
		Username: req.Username,
		Email:    req.Email,
		Password: string(hashedPassword),
	}

	result := db.Create(&user)
	if result.Error != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Could not create user",
		})
	}

	return c.Status(201).JSON(fiber.Map{
		"message": "User registered successfully",
		"user":    user,
	})
}

// LoginHandler handles user login
func LoginHandler(c *fiber.Ctx) error {
	db := config.GetDB()
	
	var req models.LoginRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": "Invalid request body",
		})
	}

	// Find user by username
	var user models.User
	if err := db.Where("username = ?", req.Username).First(&user).Error; err != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	// Compare password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return c.Status(401).JSON(fiber.Map{
			"error": "Invalid credentials",
		})
	}

	// Generate JWT token
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user.ID
	claims["username"] = user.Username
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	tokenString, err := token.SignedString([]byte(config.LoadConfig().JWTSecret))
	if err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": "Could not generate token",
		})
	}

	response := models.AuthResponse{
		Token:  tokenString,
		User:   user,
		Status: "success",
	}

	return c.JSON(response)
}

// GetUserProfileHandler returns the profile of the authenticated user
func GetUserProfileHandler(c *fiber.Ctx) error {
	user := c.Locals("user").(*models.User)
	
	return c.JSON(fiber.Map{
		"user": user,
	})
}