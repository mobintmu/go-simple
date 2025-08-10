package app

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func HealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func SlowHandler(c *gin.Context) {
	// Simulate a slow response
	time.Sleep(70 * time.Second)
	c.JSON(http.StatusOK, gin.H{
		"message": "This is a slow response",
	})
}
