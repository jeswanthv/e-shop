package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jeswanthv/e-shop/product-catalog-service/pkg/db"
	"github.com/jeswanthv/e-shop/product-catalog-service/pkg/handlers"
	"github.com/jeswanthv/e-shop/product-catalog-service/pkg/middleware"
)

func main() {
	r := gin.Default()

	// Initialize the database connection
	database := db.Init()

	// Set up the product handler
	productHandler := handlers.ProductHandler{DB: database}

	// Protected routes
	auth := r.Group("/")
	auth.Use(middleware.AuthRequired)
	{
		auth.POST("/products", productHandler.CreateProduct)
		auth.GET("/products/:id", productHandler.GetProduct)
		auth.GET("/products", productHandler.ListProducts)
		auth.PUT("/products/:id", productHandler.UpdateProduct)
		auth.DELETE("/products/:id", productHandler.DeleteProduct)
	}

	// Start the server
	r.Run(":8081")
}
