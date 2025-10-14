package main

import "go-simple/internal/app"

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Enter your JWT token in the format: Bearer <token>
func main() {
	app.NewApp().Run()
}
