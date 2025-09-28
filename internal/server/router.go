package server

import (
	"fmt"
	"go-simple/docs"
	"go-simple/internal/config"
	"go-simple/internal/health"
	"go-simple/internal/product/controller"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// RegisterRoutes is an Fx Invoke that wires up your HTTP routes
func RegisterRoutes(engine *gin.Engine,
	health *health.Health,
	cfg *config.Config,
	adminProduct *controller.AdminProduct,
	clientProduct *controller.ClientProduct) {
	log.Println("🚀 Registering routes...")
	//health
	engine.GET("/health", health.Handle)
	//Admin Product routes
	adminGroup := engine.Group("/api/v1/admin/products")
	adminProduct.RegisterRoutes(adminGroup)
	//Client Product routes
	clientGroup := engine.Group("/api/v1/products")
	clientProduct.RegisterRoutes(clientGroup)
	// Swagger
	docs.SwaggerInfo.Title = "My API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Description = "This is a sample API with Gin and Swagger."
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%d", cfg.HTTPAddress, cfg.HTTPPort)
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http"} // or {"https"} in production
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Handle 404 for unknown routes
	engine.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{
			"error":   "Not Found",
			"message": "The requested endpoint does not exist",
			"path":    c.Request.URL.Path,
		})
	})
}
