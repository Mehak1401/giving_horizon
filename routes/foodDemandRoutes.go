package routes

import (
	"giving-horizon-backend/controllers"

	"github.com/gin-gonic/gin"
)

func FoodDemandRoutes(r *gin.Engine) {
	r.POST("/demand/food", controllers.CreateFoodDemand)
	r.GET("/demand/food", controllers.GetAllFoodDemands)
}
