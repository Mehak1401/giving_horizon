package routes

import (
	"giving-horizon-backend/controllers"

	"github.com/gin-gonic/gin"
)

func FundraiserRoutes(r *gin.Engine) {
	r.POST("/fundraiser/create", controllers.CreateFundraiser)
}
