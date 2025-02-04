package main

import (
	"charity-backend/database"
	"charity-backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	database.Connect()

	routes.AuthRoutes(r)
	routes.FundraiserRoutes(r)
	routes.PaymentRoutes(r)

	r.Run(":8080")
}