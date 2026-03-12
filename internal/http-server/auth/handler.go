package auth

import (
	"log/slog"
	"net/http"

	"github.com/gcinema/core/httphelper"
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

type Response struct {
	Name     string
	Password string
}

func (handler *AuthHandler) sayHi() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		data := Response{
			Name:     "Zuuuuu",
			Password: "Psfdsaf",
		}
		httphelper.ConvertToJSON(w, data, 200)
	}
}
