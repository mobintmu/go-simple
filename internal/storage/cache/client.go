package cache

import (
	"go-simple/internal/config"

	"github.com/redis/go-redis/v9"
)

func NewClient(cfg *config.Config) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr: cfg.Redis.DSN,
		DB:   cfg.Redis.DB,
	})
}
