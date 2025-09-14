package main

import (
	"log/slog"

	"github.com/cmczk/human-ed/internal/config"
	"github.com/cmczk/human-ed/internal/storage/sqlite"
	"github.com/cmczk/human-ed/lib/logger"
)

func main() {
	cfg := config.MustLoad()
	log := logger.Setup(cfg.Env)

	storage := sqlite.New(cfg.StoragePath)
	_ = storage

	log.Info("application is starting", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")
}
