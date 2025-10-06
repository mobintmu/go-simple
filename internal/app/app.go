package app

import (
	"go-simple/internal/config"
	"go-simple/internal/db/migrate"
	"go-simple/internal/db/sqlc"
	"go-simple/internal/health"
	"go-simple/internal/pkg/logger"
	productController "go-simple/internal/product/controller"
	productService "go-simple/internal/product/service"
	"go-simple/internal/server"

	"go.uber.org/fx"
)

func NewApp() *fx.App {
	return fx.New(
		fx.Provide(
			logger.NewLogger,
			config.NewConfig,
			config.InitialDB,
			health.New,
			server.NewGinEngine,
			server.CreateHTTPServer,
			migrate.NewRunner, // migration runner
			sqlc.New,
			productController.NewAdmin,
			productController.NewClient,
			productService.New,
		),
		fx.Invoke(
			server.RegisterRoutes,
			server.StartHTTPServer,
			migrate.RunMigrations, // migration hook
			logger.RegisterLoggerLifecycle,
		),
	)
}
