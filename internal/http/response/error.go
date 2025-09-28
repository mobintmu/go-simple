package response

import (
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type ErrorLog struct {
	Status int    `json:"status"`
	Error  string `json:"error"`
	Path   string `json:"path"`
	Method string `json:"method"`
}

func JSONError(ctx *gin.Context, status int, err error) {
	// Log error as JSON
	logEntry := ErrorLog{
		Status: status,
		Error:  err.Error(),
		Path:   ctx.Request.URL.Path,
		Method: ctx.Request.Method,
	}
	logJSON, _ := json.Marshal(logEntry)
	log.Println(string(logJSON))
	// Respond to client
	ctx.AbortWithStatusJSON(status, ErrorResponse{
		Error: err.Error(),
	})
}
