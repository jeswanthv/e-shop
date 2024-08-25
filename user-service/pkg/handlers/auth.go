package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jeswanthv/e-shop/user-service/pkg/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserHandler struct {
    DB *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHandler {
    return &UserHandler{DB: db}
}

// RegisterUser handles user registration
func (h *UserHandler) RegisterUser(c *gin.Context) {
    var user models.User

    // Bind the incoming JSON payload to the user model
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    // Hash the user's password before saving to the database
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to hash password"})
        return
    }
    user.Password = string(hashedPassword)

    // Save the user to the database
    if err := h.DB.Create(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
        return
    }

    // Return the user without the password field
    c.JSON(http.StatusCreated, gin.H{
        "id":        user.ID,
        "username":  user.Username,
        "email":     user.Email,
        "created_at": user.CreatedAt,
    })
}

// DummyEndpoint handles the dummy endpoint
func (h *UserHandler) DummyEndpoint(c *gin.Context) {
    c.JSON(http.StatusOK, gin.H{
        "message": "Hello, World!",
    })
}
