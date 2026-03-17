// Package auth
package auth

import (
	"log/slog"
	"net/http"

	"github.com/gcinema/core/http-server/httpreq"
	"github.com/gcinema/core/http-server/httpres"
	"github.com/go-playground/validator/v10"
)

type AuthHandler struct {
	router    *http.ServeMux
	logger    *slog.Logger
	validator *validator.Validate
}

func NewAuthHandler(mux *http.ServeMux, logger *slog.Logger) *AuthHandler {
	handler := AuthHandler{router: mux, logger: logger, validator: validator.New()}
	return &handler
}

func (handler *AuthHandler) RegisterPaths() {
	handler.router.HandleFunc("/auth/otp/send", handler.sayHi())
	handler.logger.Debug("Auth paths was added")
}

func (handler *AuthHandler) sayHi() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		payload, err := httpreq.DecodeAndValidateBody[SendOtpRequest](&w, req, handler.validator)
		if err != nil {
			return
		}

		handler.logger.Info("Получили запрос", payload.Identifier, payload.Type)
		httpres.ConvertToJSON(w, nil, 201)
	}
}
