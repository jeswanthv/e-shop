package handlers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jeswanthv/e-shop/product-catalog-service/pkg/models"
	"gorm.io/gorm"
)

type ProductHandler struct {
    DB *gorm.DB
}

// CreateProduct handles adding a new product
func (h *ProductHandler) CreateProduct(c *gin.Context) {
    var product models.Product
    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    if err := h.DB.Create(&product).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
        return
    }

    c.JSON(http.StatusCreated, product)
}

// GetProduct handles retrieving a product by ID
func (h *ProductHandler) GetProduct(c *gin.Context) {
    var product models.Product
    id := c.Param("id")

    if err := h.DB.First(&product, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }

    c.JSON(http.StatusOK, product)
}

// ListProducts handles listing all products with search, filtering, and sorting
func (h *ProductHandler) ListProducts(c *gin.Context) {
    var products []models.Product
    query := h.DB

    // Search by name or description
    search := c.Query("search")
    if search != "" {
        query = query.Where("name ILIKE ? OR description ILIKE ?", "%"+search+"%", "%"+search+"%")
    }

    // Filter by category
    categoryID := c.Query("category_id")
    if categoryID != "" {
        query = query.Where("category_id = ?", categoryID)
    }

    // Filter by price range
    minPrice := c.Query("min_price")
    maxPrice := c.Query("max_price")
    if minPrice != "" && maxPrice != "" {
        min, _ := strconv.ParseFloat(minPrice, 64)
        max, _ := strconv.ParseFloat(maxPrice, 64)
        query = query.Where("price BETWEEN ? AND ?", min, max)
    } else if minPrice != "" {
        min, _ := strconv.ParseFloat(minPrice, 64)
        query = query.Where("price >= ?", min)
    } else if maxPrice != "" {
        max, _ := strconv.ParseFloat(maxPrice, 64)
        query = query.Where("price <= ?", max)
    }

    // Sorting
    sort := c.Query("sort")
    switch sort {
    case "price_asc":
        query = query.Order("price ASC")
    case "price_desc":
        query = query.Order("price DESC")
    case "name_asc":
        query = query.Order("name ASC")
    case "name_desc":
        query = query.Order("name DESC")
    default:
        query = query.Order("created_at DESC")
    }

    // Execute the query
    if err := query.Find(&products).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products"})
        return
    }

    c.JSON(http.StatusOK, products)
}

// UpdateProduct handles updating an existing product
func (h *ProductHandler) UpdateProduct(c *gin.Context) {
    var product models.Product
    id := c.Param("id")

    if err := h.DB.First(&product, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }

    if err := c.ShouldBindJSON(&product); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
        return
    }

    if err := h.DB.Save(&product).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product"})
        return
    }

    c.JSON(http.StatusOK, product)
}

// DeleteProduct handles deleting a product
func (h *ProductHandler) DeleteProduct(c *gin.Context) {
    id := c.Param("id")

    if err := h.DB.Delete(&models.Product{}, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}

// CheckStockAvailability checks if the product is in stock
func (h *ProductHandler) CheckStockAvailability(c *gin.Context) {
    var product models.Product
    id := c.Param("id")

    if err := h.DB.First(&product, id).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
        return
    }

    if product.Stock <= 0 {
        c.JSON(http.StatusOK, gin.H{"in_stock": false, "message": "Product out of stock"})
    } else {
        c.JSON(http.StatusOK, gin.H{"in_stock": true, "stock": product.Stock})
    }
}
