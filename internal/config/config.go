package config

import (
	"database/sql"
	"go-simple/internal/db/sqlc"
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	HTTPPort    int
	HTTPAddress string
	Database    DatabaseCfg
	ENV         string
	GRPCPort    int
}

type DatabaseCfg struct {
	DSN string
}

func NewConfig() (*Config, error) {
	v := viper.New()
	v.SetEnvPrefix("APP")
	v.SetDefault("HTTP_PORT", 4000)
	v.SetDefault("HTTP_ADDRESS", "127.0.0.1")
	v.SetDefault("DATABASE_DSN", "postgresql://user:pass@localhost:5432/database?sslmode=disable")
	v.SetDefault("GRPC_PORT", 9001)
	v.SetDefault("ENV", "development")
	v.AutomaticEnv()

	cfg := &Config{
		HTTPPort:    v.GetInt("HTTP_PORT"),
		GRPCPort:    v.GetInt("GRPC_PORT"),
		HTTPAddress: v.GetString("HTTP_ADDRESS"),
		Database: DatabaseCfg{
			DSN: v.GetString("DATABASE_DSN"),
		},
		ENV: v.GetString("ENV"),
	}
	log.Printf("✅ Loaded config") //: %+v\n", cfg)
	return cfg, nil
}

func InitialDB(cfg *Config) sqlc.DBTX {
	sql, err := sql.Open("postgres", cfg.Database.DSN)
	if err != nil {
		log.Fatal(err)
	}
	return sql
}

func (cfg *Config) IsTest() bool {
	return cfg.ENV == "test"
}

func (cfg *Config) IsDevelopment() bool {
	return cfg.ENV == "development"
}
