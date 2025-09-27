package app

import (
	"go-simple/internal/config"
	"go-simple/internal/db/migrate"
	"go-simple/internal/health"
	"go-simple/internal/server"

	"go.uber.org/fx"
)

func NewApp() *fx.App {
	return fx.New(
		fx.Provide(
			config.NewConfig,
			health.New,
			server.NewGinEngine,
			server.CreateHTTPServer,
			migrate.NewRunner, // ← migration runner
		),
		fx.Invoke(
			server.RegisterRoutes,
			server.StartHTTPServer,
			migrate.RunMigrations, // ← migration hook
		),
	)
}
