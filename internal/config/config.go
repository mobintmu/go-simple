package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

type Config struct {
	HTTPPort       int
	HTTPAddress    string
	Database       DatabaseCfg
	ENV            string
	GRPCPort       int
	JWTSecret      string
	JWTExpiryHours int
	Redis          RedisCfg
}

type DatabaseCfg struct {
	DSN string
}

type RedisCfg struct {
	DSN        string
	DB         int
	Prefix     string
	DefaultTTL int // in minute
}

func NewConfig() (*Config, error) {
	v := viper.New()
	v.SetEnvPrefix("APP")
	v.AutomaticEnv()

	// Build config
	cfg := buildConfig(v)

	// Validate configuration
	if err := ValidateConfig(cfg); err != nil {
		return nil, fmt.Errorf("⚠️ invalid configuration: %w", err)
	}
	log.Printf("✅ Loaded config: %+v\n", cfg)

	return cfg, nil
}

// buildConfig constructs the Config struct from viper values
func buildConfig(v *viper.Viper) *Config {
	return &Config{
		HTTPPort:       v.GetInt("HTTP_PORT"),
		HTTPAddress:    v.GetString("HTTP_ADDRESS"),
		GRPCPort:       v.GetInt("GRPC_PORT"),
		ENV:            v.GetString("ENV"),
		JWTSecret:      v.GetString("JWT_SECRET"),
		JWTExpiryHours: v.GetInt("JWT_EXPIRY_HOURS"),
		Database: DatabaseCfg{
			DSN: v.GetString("DATABASE_DSN"),
		},
		Redis: RedisCfg{
			DSN:        v.GetString("REDIS_DSN"),
			DB:         v.GetInt("REDIS_DB"),
			Prefix:     v.GetString("REDIS_PREFIX"),
			DefaultTTL: v.GetInt("REDIS_DEFAULT_TTL"),
		},
	}
}

func (cfg *Config) IsTest() bool {
	return cfg.ENV == "test"
}

func (cfg *Config) IsDevelopment() bool {
	return cfg.ENV == "development"
}

// IsProduction returns true if environment is production
func (cfg *Config) IsProduction() bool {
	return cfg.ENV == "production"
}

func LoadEnv() {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}
	log.Printf("📋 Loaded environment: %s\n", env)

	paths := []string{
		fmt.Sprintf(".env.%s", env),
		".env",
		filepath.Join("..", fmt.Sprintf(".env.%s", env)), // ← Parent dir
		filepath.Join("..", ".env"),
	}

	for _, path := range paths {
		if err := godotenv.Load(path); err == nil {
			log.Printf("✅ Loaded: %s\n", path)
			return
		}
	}

	log.Printf("⚠️  No .env found\n")
}
