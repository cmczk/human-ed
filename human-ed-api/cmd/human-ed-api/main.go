package main

import (
	"log/slog"

	"github.com/cmczk/human-ed/internal/config"
	"github.com/cmczk/human-ed/internal/logger"
)

func main() {
	cfg := config.MustLoad()

	log := logger.Setup(cfg.Env)

	log.Info("application is starting", slog.String("env", cfg.Env))
	log.Debug("debug messages are enabled")
}
