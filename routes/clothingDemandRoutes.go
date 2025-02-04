package routes

import (
	"giving-horizon-backend/controllers"

	"github.com/gin-gonic/gin"
)

func ClothingDemandRoutes(r *gin.Engine) {
	r.POST("/demand/clothing", controllers.CreateClothingDemand)
	r.GET("/demand/clothing", controllers.GetAllClothingDemands)
}
