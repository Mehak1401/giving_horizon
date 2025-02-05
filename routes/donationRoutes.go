package routes

import (
	"giving-horizon-backend/controllers"

	"github.com/gin-gonic/gin"
)

func DonationRoutes(r *gin.Engine) {
	// Food Donations
	r.POST("/donate/food", controllers.DonateFood)
	r.GET("/inventory/food", controllers.GetFoodInventory)

	// Clothing Donations
	r.POST("/donate/clothing", controllers.DonateClothing)
	r.GET("/inventory/clothing", controllers.GetClothingInventory)
}
