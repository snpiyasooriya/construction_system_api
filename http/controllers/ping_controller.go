package controllers

import "github.com/gin-gonic/gin"

// Ping godoc
// @Summary Ping endpoint
// @Description Simple ping endpoint to test API connectivity
// @Tags Health
// @Produce json
// @Success 200 {object} map[string]string "Pong response"
// @Router /api/ping [get]
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}

// Health godoc
// @Summary Health check endpoint
// @Description Check the health status of the API service
// @Tags Health
// @Produce json
// @Success 200 {object} map[string]string "Health status"
// @Router /api/health [get]
func Health(c *gin.Context) {
	c.JSON(200, gin.H{
		"status":  "healthy",
		"service": "construction-system-api",
	})
}
