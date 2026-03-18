// Package response
package response

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gcinema/gateway/pkg/logger"
	"go.uber.org/zap"
)

type HTTPResponseHandler struct {
	log *logger.Logger
	rw  http.ResponseWriter
}

func NewHTTPResponseHandler(log *logger.Logger, rw http.ResponseWriter) *HTTPResponseHandler {
	return &HTTPResponseHandler{
		log: log,
		rw:  rw,
	}
}

func (h *HTTPResponseHandler) PanicResponse(p any, msg string) {
	err := fmt.Errorf("unexcepted panic %v", p)

	h.log.Error(msg, zap.Error(err))
	h.rw.WriteHeader(http.StatusInternalServerError)

	response := map[string]string{
		"message": msg,
		"error":   err.Error(),
	}

	if err := json.NewEncoder(h.rw).Encode(response); err != nil {
		h.log.Error("write http response", zap.Error(err))
	}
}
