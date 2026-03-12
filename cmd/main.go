package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/gcinema/gateway/internal/config"
	"github.com/gcinema/gateway/internal/http-server/auth"
)

const (
	envLocal = "local"
	envDev   = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()

	logger := setupLogger(cfg.Env)
	logger.Info("Start application",
		slog.String("env", cfg.Env),
		slog.Any("cfg", cfg))

	router := http.NewServeMux()
	server := http.Server{
		Addr: cfg.Server.Addr,
		Handler: router,
	}

	authHandler := auth.NewAuthHandler(router, logger)
	authHandler.RegisterPaths()

	logger.Info("Server started", slog.String("Addr", cfg.Server.Addr))
	server.ListenAndServe()
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(
			slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	default:
		log = slog.New(
			slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug}),
		)
	}

	return log
}
