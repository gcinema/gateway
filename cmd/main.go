package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/gcinema/gateway/internal/config"
	"github.com/gcinema/gateway/internal/handler/auth"
	"github.com/gcinema/gateway/pkg/http/middleware"
	"github.com/gcinema/gateway/pkg/http/server"
	"github.com/gcinema/gateway/pkg/logger"
	"go.uber.org/zap"
)

func main() {
	ctx, cancel := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	cfg := config.MustLoad()

	log, err := logger.NewLogger(cfg.Logger.Level, cfg.Logger.Folder)
	if err != nil {
		fmt.Println("failed to init app logger: ", err)
		os.Exit(1)
	}

	log.Debug("Logger init")

	authHandlerHTTP := auth.NewAuthHTTPHandler(nil)
	authRoutes := authHandlerHTTP.Routes()

	apiVersionRouter := server.NewAPIVersionRouter(server.APIVersion1)
	apiVersionRouter.RegisterRoutes(authRoutes...)

	httpServer := server.NewHTTPServer(
		server.NewConfig(cfg.Server.Addr, cfg.Server.ShutdownTimeout),
		log,
		middleware.RequestID(),
		middleware.Logger(log),
		middleware.Trace(),
		middleware.Panic())

	httpServer.RegisterAPIRouters(apiVersionRouter)

	if err := httpServer.Run(ctx); err != nil {
		log.Error("run server", zap.Error(err))
	}
}
