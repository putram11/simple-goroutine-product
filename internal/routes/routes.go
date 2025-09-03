package routes

import (
	"simple-goroutine-product/internal/handlers"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

// SetupRoutes configures all routes for the application
func SetupRoutes(e *echo.Echo, productHandler *handlers.ProductHandler) {
	// Swagger documentation
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// API routes
	api := e.Group("/api/v1")

	// Product routes
	products := api.Group("/products")
	products.POST("", productHandler.CreateProduct)
	products.GET("", productHandler.GetProducts)
	products.GET("/:id", productHandler.GetProduct)
	products.PUT("/:id", productHandler.UpdateProduct)
	products.DELETE("/:id", productHandler.DeleteProduct)

	// Health check
	e.GET("/health", func(c echo.Context) error {
		return c.JSON(200, map[string]string{"status": "ok"})
	})
}
