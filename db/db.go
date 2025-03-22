package db

import (
	"fmt"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	// 🔹 ตั้งค่า DSN (Database Source Name)
	dsn := os.Getenv("DB_URL")

	// 🔹 เชื่อมต่อฐานข้อมูล
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("❌ Failed to connect to database:", err)
	}

	fmt.Println("✅ Connected to PostgreSQL!")
	DB = db
}
