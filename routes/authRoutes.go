package routes

import (
	"charity-backend/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {
	r.GET("/auth/google", controllers.GoogleLogin)
	r.GET("/auth/google/callback", controllers.GoogleCallback)
}