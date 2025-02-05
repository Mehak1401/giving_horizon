package controllers

import (
	"giving-horizon-backend/database"
	"giving-horizon-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handle Clothing Donation (Add to Inventory)
func DonateClothing(c *gin.Context) {
	var donation models.ClothingInventory
	if err := c.ShouldBindJSON(&donation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Check if item exists in inventory
	var existingItem models.ClothingInventory
	result := database.DB.Where("clothing_type = ? AND size = ?", donation.ClothingType, donation.Size).First(&existingItem)

	if result.RowsAffected > 0 {
		// Item exists, update stock
		existingItem.Stock += donation.Stock
		database.DB.Save(&existingItem)
	} else {
		// New item, create entry
		database.DB.Create(&donation)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Clothing donated successfully", "data": donation})
}

// Get Clothing Inventory (Show available clothing stock)
func GetClothingInventory(c *gin.Context) {
	var inventory []models.ClothingInventory
	database.DB.Find(&inventory)
	c.JSON(http.StatusOK, inventory)
}
