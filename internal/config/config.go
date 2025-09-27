package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	HTTPPort    int
	HTTPAddress string
	Database    DatabaseCfg
}

type DatabaseCfg struct {
	DSN string
}

func NewConfig() (*Config, error) {
	v := viper.New()
	v.SetDefault("HTTP_PORT", 4000)
	v.SetDefault("HTTP_ADDRESS", "127.0.0.1")
	v.SetDefault("DATABASE_DSN", "postgresql://user:pass@localhost:5432/database?sslmode=disable")
	v.SetEnvPrefix("APP")
	v.AutomaticEnv()

	cfg := &Config{
		HTTPPort:    v.GetInt("HTTP_PORT"),
		HTTPAddress: v.GetString("HTTP_ADDRESS"),
		Database: DatabaseCfg{
			DSN: v.GetString("DATABASE_DSN"),
		},
	}
	log.Printf("✅ Loaded config: %+v\n", cfg)
	return cfg, nil
}
