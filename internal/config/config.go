package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	HTTPPort    int
	HTTPAddress string
}

func NewConfig() (*Config, error) {
	v := viper.New()
	v.SetDefault("HTTP_PORT", 4000)
	v.SetDefault("HTTP_ADDRESS", "127.0.0.1")
	v.SetEnvPrefix("APP")
	v.AutomaticEnv()

	cfg := &Config{
		HTTPPort:    v.GetInt("HTTP_PORT"),
		HTTPAddress: v.GetString("HTTP_ADDRESS"),
	}
	log.Printf("✅ Loaded config: %+v\n", cfg)
	return cfg, nil
}
