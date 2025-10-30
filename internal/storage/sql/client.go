package sql

import (
	"database/sql"
	"go-simple/internal/config"
	"go-simple/internal/storage/sql/sqlc"
	"log"
)

func InitialDB(cfg *config.Config) sqlc.DBTX {
	sql, err := sql.Open("postgres", cfg.Database.DSN)
	if err != nil {
		log.Fatal(err)
	}
	return sql
}
