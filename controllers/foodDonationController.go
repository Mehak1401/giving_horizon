package controllers

import (
	"giving-horizon-backend/database"
	"giving-horizon-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Handle Food Donation (Add to Inventory)
func DonateFood(c *gin.Context) {
	var donation models.FoodInventory
	if err := c.ShouldBindJSON(&donation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// Check if item exists in inventory
	var existingItem models.FoodInventory
	result := database.DB.Where("item_name = ?", donation.ItemName).First(&existingItem)

	if result.RowsAffected > 0 {
		// Item exists, update stock
		existingItem.Stock += donation.Stock
		database.DB.Save(&existingItem)
	} else {
		// New item, create entry
		database.DB.Create(&donation)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Food donated successfully", "data": donation})
}

// Get Food Inventory (Show available food stock)
func GetFoodInventory(c *gin.Context) {
	var inventory []models.FoodInventory
	database.DB.Find(&inventory)
	c.JSON(http.StatusOK, inventory)
}
