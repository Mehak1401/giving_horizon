package models

import "gorm.io/gorm"

// Clothing Inventory Model
type ClothingInventory struct {
	gorm.Model
	ClothingType string `json:"clothing_type"` // Example: Jackets, Shirts, Shoes
	Size         string `json:"size"`          // Example: M, L, XL
	Stock        int    `json:"stock"`         // Available quantity
	Description  string `json:"description"`   // Additional details about the clothing
	Location     string `json:"location"`      // Where the clothing is stored
}
