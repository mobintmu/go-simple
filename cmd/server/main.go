package main

import (
	"go-simple/internal/app"
	"go-simple/internal/config"
)

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Enter your JWT token in the format: Bearer <token>
func main() {
	config.LoadEnv()

	app.NewApp().Run()
}
