package app

import (
	"log"
	"os"
	"time"

	"github.com/gin-contrib/timeout"
	"github.com/gin-gonic/gin"
)

type Application struct {
	Port   string
	Router *gin.Engine
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

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	return &Application{Port: port, Router: router}
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
