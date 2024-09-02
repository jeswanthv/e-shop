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

    // Set up the product and category handlers
    productHandler := handlers.ProductHandler{DB: database}
    categoryHandler := handlers.CategoryHandler{DB: database}

    // Protected routes
    auth := r.Group("/")
    auth.Use(middleware.AuthRequired)
    {
			auth.POST("/products", productHandler.CreateProduct)
			auth.GET("/products/:id", productHandler.GetProduct)
			auth.GET("/products", productHandler.ListProducts) // Updated for search/filter/sort
			auth.PUT("/products/:id", productHandler.UpdateProduct)
			auth.DELETE("/products/:id", productHandler.DeleteProduct)

			// Inventory management
			auth.GET("/products/:id/check-stock", productHandler.CheckStockAvailability)

			// Category routes
			auth.POST("/categories", categoryHandler.CreateCategory)
			auth.GET("/categories/:id", categoryHandler.GetCategory)
			auth.GET("/categories", categoryHandler.ListCategories)
			auth.PUT("/categories/:id", categoryHandler.UpdateCategory)
			auth.DELETE("/categories/:id", categoryHandler.DeleteCategory)
    }

    // Start the server
    r.Run(":8081")
}
