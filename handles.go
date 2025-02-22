package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

// Mock email sending handler
func sendEmailHandler(c *gin.Context) {
	var emailReq EmailRequest
	if err := c.ShouldBindJSON(&emailReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	// Simulate email sending logic (no real emails sent)
	incrementEmailCount()

	c.JSON(http.StatusOK, gin.H{
		"message": "Email request received successfully",
		"status":  "Queued",
	})
}

// Get email statistics
func getStatsHandler(c *gin.Context) {
	stats := getEmailStats()
	c.JSON(http.StatusOK, stats)
}
