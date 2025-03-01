package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func sendEmailHandler(c *gin.Context) {
	var req EmailRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	if emailStats.EmailsPerUser[req.Sender] >= emailStats.DailyLimit {
		c.JSON(http.StatusForbidden, gin.H{"error": "Daily limit reached for sender"})
		return
	}

	incrementEmailCount(req.Sender)

	c.JSON(http.StatusOK, gin.H{"message": "Email accepted for delivery (mock)", "sender": req.Sender})
}

func getStatsHandler(c *gin.Context) {
	stats := getEmailStats()
	c.JSON(http.StatusOK, stats)
}
