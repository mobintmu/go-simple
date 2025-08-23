package health

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type Health struct {
}

func New() *Health {
	return &Health{}
}

// @Summary Get health status
// @Description Returns the health status of the API
// @Tags Health
// @Produce json
// @Success 200 {object} HealthResponse "OK"
// @Router /health [get]
func (h *Health) HealthHandler(c *gin.Context) {
	c.JSON(http.StatusOK, HealthResponse{Message: "OK"})
}

type HealthResponse struct {
	Message string `json:"message" example:"OK"`
}

// @Summary Simulate a slow endpoint
// @Description Simulates a slow response to test timeout handling
// @Tags Health
// @Produce json
// @Success 200 {object} HealthResponse "This is a slow response"
// @Router /slow [get]
func (h *Health) SlowHandler(c *gin.Context) {
	// Simulate a slow response
	time.Sleep(70 * time.Second)
	c.JSON(http.StatusOK, gin.H{
		"message": "This is a slow response",
	})
}
