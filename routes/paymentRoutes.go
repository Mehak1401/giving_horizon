package routes

import (
	"charity-backend/controllers"
	"github.com/gin-gonic/gin"
)

func PaymentRoutes(r *gin.Engine) {
	r.POST("/donate", controllers.ProcessDonation)
}