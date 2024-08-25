package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jeswanthv/e-shop/user-service/pkg/db"
	"github.com/jeswanthv/e-shop/user-service/pkg/handlers"
)

func main() {
    r := gin.Default()

    // Initialize the database
    database := db.Init()

    // Set up user handler
    userHandler := handlers.NewUserHandler(database)

    // Define routes
    r.POST("/register", userHandler.RegisterUser)
    r.GET("/dummy", userHandler.DummyEndpoint) // Dummy endpoint

    // Start the server
    r.Run(":8080")
}
