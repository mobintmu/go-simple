package migrate

import (
	"go-simple/internal/config"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Runner struct {
	DSN string
}

// internal/db/migrate/runner.go
func NewRunner(cfg *config.Config) *Runner {
	return &Runner{DSN: cfg.Database.DSN}
}

func (r *Runner) Run() {
	m, err := migrate.New(
		"file://internal/storage/sql/migrations",
		r.DSN,
	)
	if err != nil {
		log.Fatalf("migration init error: %v", err)
	}
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("migration run error: %v", err)
	}
	log.Println("✅ Migrations applied successfully")
}
