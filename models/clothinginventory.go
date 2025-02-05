package models

import "gorm.io/gorm"

type ClothingInventory struct {
	gorm.Model
	ClothingType string `json:"clothing_type"`
	Size         string `json:"size"`
	Stock        int    `json:"stock"`
	Description  string `json:"description"`
	Location     string `json:"location"`
}
