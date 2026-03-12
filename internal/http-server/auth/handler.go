package auth

import (
	"log/slog"
	"net/http"
)

type AuthHandler struct {
	router *http.ServeMux
	logger *slog.Logger
}

func NewAuthHandler(mux *http.ServeMux, logger *slog.Logger) *AuthHandler {
	handler := AuthHandler{router: mux, logger: logger}
	return &handler
}

func (handler *AuthHandler) RegisterPaths() {
	handler.router.HandleFunc("/say-hi", handler.sayHi())
	handler.logger.Debug("Auth paths was added")
}

func (handler *AuthHandler) sayHi() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		handler.logger.Info("Hello bro")
	}
}
