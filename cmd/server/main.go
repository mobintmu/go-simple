package main

import (
	application "go-simple/cmd/server/app"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	app := application.New()
	app.RegisterRoutes()
	app.StartServer()

	// Wait for interrupt signal to gracefully shut down
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig // Blocks until signal is received
}
