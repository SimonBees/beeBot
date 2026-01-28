package config

import (
	"log"
	"gorm.io/driver/sqlite"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"github.com/SimonBees/beeBot/models"
)

var db *gorm.DB

// InitDatabase initializes the database connection
func InitDatabase() {
	cfg := LoadConfig()

	var err error
	logLevel := logger.Silent
	if cfg.Env == "development" {
		logLevel = logger.Info
	}

	// Use SQLite for development if no DB password is set
	if cfg.DBPass == "" || cfg.DBHost == "localhost" {
		log.Println("Using SQLite database for development")
		db, err = gorm.Open(sqlite.Open("beebot.db"), &gorm.Config{
			Logger: logger.Default.LogMode(logLevel),
		})
	} else {
		// Use PostgreSQL for production
		dsn := "host=" + cfg.DBHost +
			" port=" + cfg.DBPort +
			" user=" + cfg.DBUser +
			" password=" + cfg.DBPass +
			" dbname=" + cfg.DBName +
			" sslmode=disable"

		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
			Logger: logger.Default.LogMode(logLevel),
		})
	}

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Database connected successfully")

	// Auto migrate tables
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}

	log.Println("Database migration completed")
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
	return db
}
