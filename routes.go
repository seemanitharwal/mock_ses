package main

import "github.com/gin-gonic/gin"

func setupRoutes(router *gin.Engine) {
	router.POST("/send-email", sendEmailHandler)
	router.GET("/stats", getStatsHandler)
}
