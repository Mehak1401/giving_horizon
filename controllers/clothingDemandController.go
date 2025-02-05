package controllers

import (
	"giving-horizon-backend/database"
	"giving-horizon-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateClothingDemand(c *gin.Context) {
	var demand models.ClothingDemand
	if err := c.ShouldBindJSON(&demand); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	var inventory models.ClothingInventory
	result := database.DB.Where("clothing_type = ? AND size = ?", demand.ClothingType, demand.Size).First(&inventory)

	if result.RowsAffected == 0 || inventory.Stock < demand.Quantity {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Not enough stock available"})
		return
	}

	inventory.Stock -= demand.Quantity
	database.DB.Save(&inventory)

	database.DB.Create(&demand)
	c.JSON(http.StatusOK, gin.H{"message": "Clothing demand created and inventory updated", "data": demand})
}
