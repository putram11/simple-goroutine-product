package main

import (
	"log"
	"simple-goroutine-product/internal/database"

	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Connect to database
	database.ConnectDatabase()

	log.Println("Database migration completed successfully")
}
