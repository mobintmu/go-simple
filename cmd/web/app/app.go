package application

import (
	"log"
	"net/http"
	"os"

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

	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	return &Application{Port: port, Router: router}
}

func (app *Application) StartServer() {
	//start the server
	go func() {
		log.Printf("Starting Gin server at port %s", app.Port)
		err := app.Router.Run(app.Port)
		if err != nil {
			log.Fatal(err)
		}
	}()
}

func (a *Application) Routes() {
	a.Router.GET("/", helloHandler)
}

func helloHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello world",
	})
}
