package main

import (
	"giving-horizon-backend/database"
	"giving-horizon-backend/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.Use(CORSMiddleware())
	database.Connect()

	routes.AuthRoutes(r)
	routes.FundraiserRoutes(r)
	routes.PaymentRoutes(r)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	if err := r.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}
