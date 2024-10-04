package database

import (
	"go-fiber-api/internal/models"
	"log"

	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {

	// Handle error env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Mengambil value dari ENV
	host := os.Getenv("DB_PG_HOST")
	user := os.Getenv("DB_PG_USER")
	password := os.Getenv("DB_PG_PASSWORD")
	dbname := os.Getenv("DB_PG_NAME")
	port := os.Getenv("DB_PG_PORT")

	// String koneksi ke database PostgreSQL
	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=disable TimeZone=Asia/Jakarta"

	// Hardcode
	// dsn := "host=localhost user=pgadmin password=P@ssw0rd# dbname=postgres port=3000 sslmode=disable TimeZone=Asia/Jakarta"

	// var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}

	// Migrasi skema database
	DB.AutoMigrate(&models.User{})
}
