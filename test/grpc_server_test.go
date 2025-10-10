package test

import (
	"context"
	app "go-simple/internal/app"
	"os"
	"testing"
	"time"

	"go.uber.org/fx"
)

func StartGRPCServer() *fx.App {
	os.Setenv("APP_GRPC_PORT", "9090")
	os.Setenv("APP_ENV", "test")
	os.Setenv("APP_DATABASE_DSN", "postgresql://user:pass@localhost:5432/database?sslmode=disable")

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
