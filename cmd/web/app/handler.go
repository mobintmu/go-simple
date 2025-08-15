package app

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// @Summary Get health status
// @Description Returns the health status of the API
// @Tags Health
// @Produce json
// @Success 200 {object} HealthResponse "OK"
// @Router /health [get]
func HealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, HealthResponse{Message: "OK"})
}

type HealthResponse struct {
	Message string `json:"message" example:"OK"`
}

func SlowHandler(c *gin.Context) {
	// Simulate a slow response
	time.Sleep(70 * time.Second)
	c.JSON(http.StatusOK, gin.H{
		"message": "This is a slow response",
	})
}
