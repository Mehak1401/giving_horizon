package main

import (
	"fmt"
	"giving-horizon-backend/database"
	"giving-horizon-backend/routes"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database
	database.Connect()

	// Create a new Gin router
	r := gin.Default()

	// Load routes
	routes.AuthRoutes(r)
	routes.FundraiserRoutes(r)
	routes.PaymentRoutes(r)

	// Fetch the Railway port or use 8080 for local dev
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port for local development
	}

	// Start the server
	fmt.Println("🚀 Server running on port:", port)
	r.Run("0.0.0.0" + port)
}
