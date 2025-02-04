package controllers

import (
	"giving-horizon-backend/database"
	"giving-horizon-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ProcessDonation(c *gin.Context) {
	var donation models.Donation
	if err := c.BindJSON(&donation); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	database.DB.Create(&donation)
	c.JSON(http.StatusOK, gin.H{"message": "Donation Successful", "donation": donation})
}
