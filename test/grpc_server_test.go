package test

import (
	"context"
	app "go-simple/internal/app"
	"go-simple/internal/config"
	"os"
	"testing"
	"time"

	"go.uber.org/fx"
)

func StartGRPCServer() *fx.App {
	// Change to project root
	os.Chdir("..")
	defer os.Chdir("test")
	os.Setenv("APP_ENV", "test")
	config.LoadEnv()

	a := app.NewApp()
	go a.Run()
	time.Sleep(300 * time.Millisecond) // give server time to start
	return a
}

func WithGRPCTestServer(t *testing.T, testFunc func()) {
	a := StartGRPCServer()
	defer a.Stop(context.Background())
	testFunc()
}
