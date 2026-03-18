package middleware

import (
	"context"
	"net/http"

	"github.com/gcinema/gateway/pkg/ctxkey"
	"github.com/gcinema/gateway/pkg/logger"
	"go.uber.org/zap"
)

func Logger(log *logger.Logger) Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			requestID := r.Header.Get(requestIDHeader)

			l := log.With(
				zap.String("request_id", requestID),
				zap.String("url", r.URL.String()),
			)

			ctx := context.WithValue(r.Context(), ctxkey.Log, l)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
