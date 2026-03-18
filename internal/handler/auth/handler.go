// Package auth
package auth

import (
	"net/http"

	"github.com/gcinema/gateway/internal/handler/auth/dto"
	"github.com/gcinema/gateway/pkg/http/httpreq"
	"github.com/gcinema/gateway/pkg/http/httpres"
	"github.com/gcinema/gateway/pkg/http/server"
	"github.com/gcinema/gateway/pkg/logger"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

type AuthHTTPHandler struct {
	authService AuthService
	validator   *validator.Validate
}

type AuthService interface{}

func NewAuthHTTPHandler(authService AuthService) *AuthHTTPHandler {
	return &AuthHTTPHandler{
		authService: authService,
		validator:   validator.New(),
	}
}

func (h *AuthHTTPHandler) Routes() []server.Route {
	return []server.Route{
		{
			Method:  http.MethodPost,
			Path:    "/auth/otp/send",
			Handler: h.SendOTP,
		},
	}
}

func (h *AuthHTTPHandler) SendOTP(w http.ResponseWriter, r *http.Request) {
	log := logger.FromContext(r.Context())

	respHandler := httpres.NewHTTPResponseHandler(log, w)

	var req dto.SendOtpRequest
	if err := httpreq.DecodeAndValidateBody(r, h.validator, &req); err != nil {
		respHandler.ErrorResponse(err, "failed to decode and validate HTTP req")
		return
	}

	log.Debug("Делаем какую то обработку над данными",
		zap.String("Identifier", req.Identifier),
		zap.String("Type", string(req.Type)))

	w.WriteHeader(http.StatusOK)
}
