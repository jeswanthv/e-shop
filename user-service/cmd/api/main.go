package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jeswanthv/e-shop/user-service/pkg/db"
	"github.com/jeswanthv/e-shop/user-service/pkg/handlers"
	"github.com/jeswanthv/e-shop/user-service/pkg/middleware"
)

func main() {
    r := gin.Default()

    // Initialize the database
    database := db.Init()

    // Set up user handler
    userHandler := handlers.NewUserHandler(database)

    // Define routes
    r.POST("/register", userHandler.RegisterUser)
    r.POST("/login", userHandler.LoginUser)

    // Password reset routes
	r.POST("/request-reset", userHandler.RequestPasswordReset)
	r.POST("/reset-password", userHandler.ResetPassword)
    
    // Protected routes
	auth := r.Group("/")
	auth.Use(middleware.AuthRequired)
	{
		auth.PUT("/profile", userHandler.UpdateProfile)
	}

    // Start the server
    r.Run(":8080")
}
