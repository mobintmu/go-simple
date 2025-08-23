package app

import (
	health "go-simple/internal/health/controller"
	"log"
	"os"
	"time"

	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
)

type Application struct {
	Port       string
	Router     *gin.Engine
	healthCtrl *health.Health
}

const defaultPort = ":4000"

func New() *Application {
	// Set release mode
	gin.SetMode(gin.DebugMode)

	// Use New() for full control, or Default() if you're okay with built-ins
	router := gin.New()

	// Add middleware explicitly
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Use(timeout.New(timeout.WithTimeout(60 * time.Second))) // timeout middleware

	// Register Controller
	healthCtrl := health.New()

	//set port
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	return &Application{
		Port:       port,
		Router:     router,
		healthCtrl: healthCtrl,
	}
}

func (app *Application) StartServer() {
	go func() {
		log.Printf("Starting Gin server at port %s", app.Port)
		err := app.Router.Run(app.Port)
		if err != nil {
			log.Fatal(err)
		}
	}()
}
