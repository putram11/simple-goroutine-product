// @title Simple Product API
// @version 1.0
// @description A simple CRUD API for products using Go, Echo, GORM with Goroutines
// @host localhost:8080
// @BasePath /
package main

import (
	"log"
	"os"
	"simple-goroutine-product/internal/database"
	"simple-goroutine-product/internal/handlers"
	"simple-goroutine-product/internal/presenters"
	"simple-goroutine-product/internal/repositories"
	"simple-goroutine-product/internal/routes"
	"simple-goroutine-product/internal/validators"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	// Connect to database
	database.ConnectDatabase()

	// Initialize repositories
	productRepo := repositories.NewProductRepository(database.GetDB())

	// Initialize presenters
	productPresenter := presenters.NewProductPresenter(productRepo)

	// Initialize handlers
	productHandler := handlers.NewProductHandler(productPresenter)

	// Initialize Echo
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Custom validator
	e.Validator = validators.NewValidator()

	// Setup routes
	routes.SetupRoutes(e, productHandler)

	// Get port from environment
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = "8080"
	}

	// Start server
	log.Printf("Server starting on port %s", port)
	log.Fatal(e.Start(":" + port))
}
