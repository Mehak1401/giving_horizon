package controllers

import (
	"giving-horizon-backend/database"
	"giving-horizon-backend/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func CreateFundraiser(c *gin.Context) {
	var fundraiser models.Fundraiser
	if err := c.BindJSON(&fundraiser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	database.DB.Create(&fundraiser)
	c.JSON(http.StatusOK, fundraiser)
}
