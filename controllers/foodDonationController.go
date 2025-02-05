package controllers

import (
	"giving-horizon-backend/database"
	"giving-horizon-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DonateFood(c *gin.Context) {
	var donation models.FoodInventory
	if err := c.ShouldBindJSON(&donation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var existingItem models.FoodInventory
	result := database.DB.Where("item_name = ?", donation.ItemName).First(&existingItem)

	if result.RowsAffected > 0 {
		existingItem.Stock += donation.Stock
		database.DB.Save(&existingItem)
	} else {
		database.DB.Create(&donation)
	}

	c.JSON(http.StatusOK, gin.H{"message": "Food donated successfully", "data": donation})
}

func GetFoodInventory(c *gin.Context) {
	var inventory []models.FoodInventory
	database.DB.Find(&inventory)
	c.JSON(http.StatusOK, inventory)
}
