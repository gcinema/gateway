// Package auth
package auth

import (
	"net/http"

	"github.com/gcinema/core/http-server/httpreq"
	"github.com/gcinema/core/http-server/httpres"
	"github.com/gcinema/gateway/internal/handler/auth/dto"
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

func (h *AuthHTTPHandler) SendOTP(w http.ResponseWriter, req *http.Request) {
	log := logger.FromContext(req.Context())
	payload, err := httpreq.DecodeAndValidateBody[dto.SendOtpRequest](&w, req, h.validator)
	if err != nil {
		return
	}

	log.Debug("Делаем какую то обработку над данными",
		zap.String("Identifier", payload.Identifier),
		zap.String("Type", string(payload.Type)))
	httpres.ConvertToJSON(w, nil, http.StatusOK)
}
