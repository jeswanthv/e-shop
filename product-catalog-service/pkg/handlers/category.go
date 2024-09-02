package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jeswanthv/e-shop/product-catalog-service/pkg/models"
	"gorm.io/gorm"
)

type CategoryHandler struct {
    DB *gorm.DB
}

// CreateCategory handles adding a new category
func (h *CategoryHandler) CreateCategory(c *gin.Context) {
	var category models.Category
	if err := c.ShouldBindJSON(&category); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
	}

	// Check if a category with the same name already exists
	var existingCategory models.Category
	if err := h.DB.Where("name = ?", category.Name).First(&existingCategory).Error; err == nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Category already exists"})
			return
	}

	// If the category doesn't exist, create it
	if err := h.DB.Create(&category).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category"})
			return
	}

	c.JSON(http.StatusCreated, category)
}

// GetCategory handles retrieving a category by ID
func (h *CategoryHandler) GetCategory(c *gin.Context) {
    var category models.Category
    id := c.Param("id")

    if err := h.DB.First(&category, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
        return
    }

    c.JSON(http.StatusOK, category)
}

// ListCategories handles listing all categories
func (h *CategoryHandler) ListCategories(c *gin.Context) {
    var categories []models.Category

    if err := h.DB.Find(&categories).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve categories"})
        return
    }

    c.JSON(http.StatusOK, categories)
}

// UpdateCategory handles updating an existing category
func (h *CategoryHandler) UpdateCategory(c *gin.Context) {
	var updatedCategory models.Category
	id := c.Param("id")

	// Bind the JSON body to the updatedCategory struct
	if err := c.ShouldBindJSON(&updatedCategory); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
			return
	}

	// Find the existing category by ID
	var existingCategory models.Category
	if err := h.DB.First(&existingCategory, id).Error; err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
			return
	}

	// Check if the new name already exists (and it's not the current category)
	if existingCategory.Name != updatedCategory.Name {
			var duplicateCategory models.Category
			if err := h.DB.Where("name = ?", updatedCategory.Name).First(&duplicateCategory).Error; err == nil {
					c.JSON(http.StatusBadRequest, gin.H{"error": "Category name already in use"})
					return
			}
	}

	// Update the category fields
	existingCategory.Name = updatedCategory.Name

	// Save the updated category
	if err := h.DB.Save(&existingCategory).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update category"})
			return
	}

	c.JSON(http.StatusOK, existingCategory)
}

// DeleteCategory handles deleting a category
func (h *CategoryHandler) DeleteCategory(c *gin.Context) {
    id := c.Param("id")

    if err := h.DB.Delete(&models.Category{}, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Category deleted"})
}
