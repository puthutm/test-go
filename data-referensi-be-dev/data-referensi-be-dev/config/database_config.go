package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	var err error

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	username := os.Getenv("DB_USERNAME")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	database := os.Getenv("DB_NAME")

	if host == "" {
		host = "localhost"
	}
	if port == "" {
		port = "5432"
	}
	if username == "" {
		username = "postgres"
	}
	if database == "" {
		database = "data_referensi_db"
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta", host, username, password, database, port)

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Printf("⚠️ Gagal konek ke PostgreSQL (%s:%s): %v", host, port, err)
		log.Println("💡 Pastikan layanan PostgreSQL service/Docker sudah dinyalakan pada port 5432.")
	} else {
		log.Println("✅ PostgreSQL Database connected successfully!")
	}
}
