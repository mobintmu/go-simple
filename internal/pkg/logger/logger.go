package logger

import (
	"context"
	"strings"

	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewLogger() (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{"stdout"}
	return cfg.Build()
}

func RegisterLoggerLifecycle(lc fx.Lifecycle, log *zap.Logger) {
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			log.Info("⏹️ On Stop Logger Life cycle ⏹️")
			if err := log.Sync(); err != nil && !isIgnorableSyncError(err) {
				return err
			}
			return nil
		},
	})
}

func isIgnorableSyncError(err error) bool {
	return strings.Contains(err.Error(), "invalid argument") ||
		strings.Contains(err.Error(), "bad file descriptor")
}
