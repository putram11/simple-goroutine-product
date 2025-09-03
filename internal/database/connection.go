package database

import (
	"fmt"
	"log"
	"os"
	"simple-goroutine-product/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// DB holds the database connection
var DB *gorm.DB

// ConnectDatabase initializes the database connection
func ConnectDatabase() {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	database, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Auto migrate the schema
	err = database.AutoMigrate(&models.Product{})
	if err != nil {
		log.Fatal("Failed to migrate database:", err)
	}

	DB = database
	log.Println("Database connected successfully")
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
	return DB
}
