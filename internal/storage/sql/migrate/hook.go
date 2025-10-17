package migrate

import "go-simple/internal/config"

func RunMigrations(runner *Runner, cfg *config.Config) {
	if cfg.IsTest() {
		return
	}
	runner.Run()
}
