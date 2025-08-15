package app

import (
	"go-simple/docs"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (a *Application) Routes() {
	api := a.Router.Group("/api/v1/")
	api.GET("/health", HealthHandler)
	api.GET("/slow", SlowHandler)

	//  Set Swagger metadata dynamically
	docs.SwaggerInfo.Title = "My API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Description = "This is a sample API with Gin and Swagger."
	docs.SwaggerInfo.Host = "127.0.0.1:4000"
	docs.SwaggerInfo.BasePath = "/api/v1"
	docs.SwaggerInfo.Schemes = []string{"http"} // or {"https"} in production
	a.Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
