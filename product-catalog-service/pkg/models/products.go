package models

import (
	"time"
)

type Product struct {
    ID          uint      `json:"id" gorm:"primaryKey"`
    Name        string    `json:"name" gorm:"not null"`
    Description string    `json:"description"`
    Price       float64   `json:"price" gorm:"not null"`
    CategoryID  uint      `json:"category_id"`
    Stock       int       `json:"stock" gorm:"not null"`
    CreatedAt   time.Time `json:"created_at"`
    UpdatedAt   time.Time `json:"updated_at"`
}
