package models

import "gorm.io/gorm"

type FoodInventory struct {
	gorm.Model
	ItemName    string `json:"item_name"`
	Stock       int    `json:"stock"`
	Description string `json:"description"`
	Location    string `json:"location"`
}
