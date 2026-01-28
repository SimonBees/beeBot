package config

import (
	"os"
)

type Config struct {
	Port     string
	Env      string
	DBHost   string
	DBPort   string
	DBUser   string
	DBPass   string
	DBName   string
	JWTSecret string
}

func LoadConfig() Config {
	return Config{
		Port: getEnv("PORT", "3000"),
		Env:  getEnv("ENV", "development"),
		DBHost: getEnv("DB_HOST", "localhost"),
		DBPort: getEnv("DB_PORT", "5432"),
		DBUser: getEnv("DB_USER", "postgres"),
		DBPass: getEnv("DB_PASS", ""),
		DBName: getEnv("DB_NAME", "beebot"),
		JWTSecret: getEnv("JWT_SECRET", "default_secret_key"),
	}
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}