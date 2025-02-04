package controllers

import (
	"giving-horizon-backend/database"
	"giving-horizon-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Create Food Demand
func CreateFoodDemand(c *gin.Context) {
	var demand models.FoodDemand
	if err := c.ShouldBindJSON(&demand); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	database.DB.Create(&demand)
	c.JSON(http.StatusOK, gin.H{"message": "Food demand created successfully", "data": demand})
}

// Get All Food Demands
func GetAllFoodDemands(c *gin.Context) {
	var demands []models.FoodDemand
	database.DB.Find(&demands)
	c.JSON(http.StatusOK, demands)
}
