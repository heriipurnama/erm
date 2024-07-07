package config

import (
	"log"
	"os"

	"dbo/erm/models" // Adjust the import path as necessary

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDatabase() (*gorm.DB, error) {
	// Load environment variables from .env file
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error loading .env file")
	}

	// Retrieve the DSN from the environment variable
	dsn := os.Getenv("DB_DSN")
	if dsn == "" {
		log.Fatal("DB_DSN environment variable is not set")
	}

	// Connect to the database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
		return nil, err
	}

	// Automigrate models
	if err := db.AutoMigrate(&models.Customer{}, &models.Order{}, &models.Authentication{}); err != nil {
		log.Fatalf("Could not migrate database: %v", err)
		return nil, err
	}

	DB = db
	return db, nil
}
