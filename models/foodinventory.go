package models

import "gorm.io/gorm"

// Food Inventory Model
type FoodInventory struct {
	gorm.Model
	ItemName    string `json:"item_name"`   // Example: Rice, Bread, Milk
	Stock       int    `json:"stock"`       // Available quantity
	Description string `json:"description"` // Additional details about the food
	Location    string `json:"location"`    // Where the food is stored
}
