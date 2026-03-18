// Package server
package server

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/gcinema/gateway/pkg/http/middleware"
	"github.com/gcinema/gateway/pkg/logger"
	"go.uber.org/zap"
)

type HTTPServer struct {
	mux         *http.ServeMux
	config      Config
	log         *logger.Logger
	middlewares []middleware.Middleware
}

func NewHTTPServer(
	cfg Config,
	log *logger.Logger,
	middlewares ...middleware.Middleware,
) *HTTPServer {
	return &HTTPServer{
		mux:         http.NewServeMux(),
		config:      cfg,
		log:         log,
		middlewares: middlewares,
	}
}

func (h *HTTPServer) RegisterAPIRouters(routers ...*APIVersionRouter) {
	for _, router := range routers {
		prefix := "/api/" + string(router.apiVersion)

		h.mux.Handle(
			prefix+"/",
			http.StripPrefix(prefix, router))
	}
}

func (h *HTTPServer) Run(ctx context.Context) error {
	mux := middleware.ChainMiddleware(h.mux, h.middlewares...)
	server := &http.Server{
		Addr:    h.config.Addr,
		Handler: mux,
	}

	ch := make(chan error, 1)
	go func() {
		defer close(ch)

		h.log.Warn("start HTTP server", zap.String("addr", h.config.Addr))

		err := server.ListenAndServe()
		if !errors.Is(err, http.ErrServerClosed) {
			ch <- err
		}
	}()

	select {
	case err := <-ch:
		if err != nil {
			return fmt.Errorf("start ListenAndServe", err)
		}
	case <-ctx.Done():
		h.log.Warn("shutdown HTTP server")

		shutDownCtx, cancel := context.WithTimeout(
			context.Background(),
			h.config.ShutdownTimeout)
		defer cancel()

		if err := server.Shutdown(shutDownCtx); err != nil {
			_ = server.Close()

			return fmt.Errorf("shutdown HTTP server", err)
		}

		h.log.Warn("HTTP server stoped")
	}

	return nil
}
