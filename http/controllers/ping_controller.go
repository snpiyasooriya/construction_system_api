package controllers

import "github.com/gin-gonic/gin"

func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

func Health(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  "healthy",
		"service": "construction-system-api",
	})
}
