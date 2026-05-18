package core_logger

import (
	"fmt"
	"log/slog"
	"os"

	"github.com/kelseyhightower/envconfig"
)

type loggerConfig struct {
	Env string `envconfig:"env" default:"dev"`
}

func MustSetupLogger() *slog.Logger {
	var log *slog.Logger
	var cfg loggerConfig

	err := envconfig.Process("", &cfg)
	if err != nil {
		err = fmt.Errorf("load logger config: %w", err)
		panic(err)
	}


	switch cfg.Env {
	case "local": 
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case "dev":
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}))
	case "prod":
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	log.Info("logger initialized with env", slog.String("env", cfg.Env))

	return log
}