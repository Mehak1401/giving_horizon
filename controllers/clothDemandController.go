package controllers

import (
	"giving-horizon-backend/database"
	"giving-horizon-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create Clothing Demand
func CreateClothingDemand(c *gin.Context) {
	var demand models.ClothingDemand
	if err := c.ShouldBindJSON(&demand); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	database.DB.Create(&demand)
	c.JSON(http.StatusOK, gin.H{"message": "Clothing demand created successfully", "data": demand})
}

// Get All Clothing Demands
func GetAllClothingDemands(c *gin.Context) {
	var demands []models.ClothingDemand
	database.DB.Find(&demands)
	c.JSON(http.StatusOK, demands)
}
