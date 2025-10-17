package app

import (
	"go-simple/internal/config"
	"go-simple/internal/health"
	"go-simple/internal/pkg/logger"
	productController "go-simple/internal/product/controller"
	productService "go-simple/internal/product/service"
	"go-simple/internal/server"
	"go-simple/internal/storage/cache"
	"go-simple/internal/storage/sql/migrate"
	"go-simple/internal/storage/sql/sqlc"

	"go.uber.org/fx"
)

func NewApp() *fx.App {
	return fx.New(
		fx.Provide(
			logger.NewLogger,
			config.NewConfig,
			config.InitialDB,
			//server
			health.New,
			server.NewGinEngine,
			server.CreateHTTPServer,
			server.CreateGRPCServer,
			//db
			migrate.NewRunner, // migration runner
			sqlc.New,
			//cache
			cache.NewClient,
			cache.NewCacheStore,
			//controller
			productController.NewAdmin,
			productController.NewClient,
			productController.NewGRPC,
			//service
			productService.New,
		),
		fx.Invoke(
			server.RegisterRoutes,
			server.StartHTTPServer,
			server.StartGRPCServer,
			//migration
			migrate.RunMigrations,
			//life cycle
			logger.RegisterLoggerLifecycle,
			server.GRPCLifeCycle,
		),
	)
}
