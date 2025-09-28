package health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Health struct {
}

func New() *Health {
	return &Health{}
}

// Health godoc
// @Summary Get health status
// @Description Returns the health status of the API
// @Tags Health
// @Produce json
// @Success 200 {object} HealthResponse "OK"
// @Router /health [get]
func (h *Health) Handle(c *gin.Context) {
	c.JSON(http.StatusOK, HealthResponse{Message: "OK"})
}

type HealthResponse struct {
	Message string `json:"message" example:"OK"`
}
