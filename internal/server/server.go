package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"go-simple/docs"
	"go-simple/internal/config"
	"go-simple/internal/health"

	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/fx"
)

func NewGinEngine() *gin.Engine {
	if gin.Mode() != gin.ReleaseMode {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()
	r.Use(gin.Logger(),
		gin.Recovery(),
		timeout.New(timeout.WithTimeout(60*time.Second)))

	return r
}

func CreateHTTPServer(engine *gin.Engine, cfg *config.Config) *http.Server {
	return &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.HTTPPort),
		Handler: engine,
	}
}

// NewHTTPServer is an Fx provider that sets up the HTTP server in the Fx lifecycle
func StartHTTPServer(lc fx.Lifecycle, srv *http.Server) {
	lc.Append(fx.Hook{
		OnStart: func(startCtx context.Context) error {
			log.Printf("🚀 HTTP server starting on %s", srv.Addr)
			go func() {
				if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
					log.Fatalf("server error: %v", err)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("🛑 Shutting down HTTP server...")
			shutdownCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
			defer cancel()
			if err := srv.Shutdown(shutdownCtx); err != nil {
				log.Printf("Shutdown error: %v", err)
				return err
			}
			log.Println("✅ Server shutdown complete.")
			return nil
		},
	})
}

// RegisterRoutes is an Fx Invoke that wires up your HTTP routes
func RegisterRoutes(engine *gin.Engine, health *health.Health, cfg *config.Config) {
	log.Println("🚀 Registering routes...")
	engine.GET("/health", health.Handle)
	//  Set Swagger metadata dynamically
	docs.SwaggerInfo.Title = "My API"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Description = "This is a sample API with Gin and Swagger."
	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%d", cfg.HTTPAddress, cfg.HTTPPort)
	docs.SwaggerInfo.BasePath = "/"
	docs.SwaggerInfo.Schemes = []string{"http"} // or {"https"} in production
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
