package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"api/models"
)

var DB *gorm.DB

func ConnectDB() {
	// Loading environment variables for local development (optional)
	err := godotenv.Load()
	if err != nil {
		log.Println("⚠️ Warning: No .env file found")
	}

	// Check if DATABASE_URL is set (Railway environment variable)
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("❌ DATABASE_URL is not set")
	}

	// Open the database connection using the full DATABASE_URL
	db, err := gorm.Open(postgres.Open(databaseURL), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
	}

	fmt.Println("✅ Database Connected Successfully!")

	// Perform AutoMigrate for your models
	err = db.AutoMigrate(
		&models.Brand{},
		&models.Model{},
		&models.Phone{},
		&models.Condition{},
		&models.PricingAdjustment{},
	)
	if err != nil {
		log.Fatal("❌ AutoMigrate failed:", err)
	} else {
		fmt.Println("✅ AutoMigrate completed!")
	}
	fmt.Println("✅ Database Migrated Successfully!")

	DB = db
}
