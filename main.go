package main

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Initialize routes
	setupRoutes(router)

	// Start server
	log.Println("Mock AWS SES API running on port 8080...")
	http.ListenAndServe(":8080", router)
}
