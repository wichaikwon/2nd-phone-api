package config

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"api/models"
)

var DB *gorm.DB

func ConnectDB() {
	// ดึงค่าจาก DATABASE_URL ที่ Railway ตั้งให้
	databaseURL := os.Getenv("DATABASE_URL")
	if databaseURL == "" {
		log.Fatal("❌ DATABASE_URL is not set")
	}

	// เชื่อมต่อกับฐานข้อมูล PostgreSQL บน Railway
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

	DB = db
}
