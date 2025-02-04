package routes

import (
	"charity-backend/controllers"
	"github.com/gin-gonic/gin"
)

func FundraiserRoutes(r *gin.Engine) {
	r.POST("/fundraiser/create", controllers.CreateFundraiser)
}