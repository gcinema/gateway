package middleware

import (
	"net/http"
	"time"

	"github.com/gcinema/gateway/pkg/http/response"
	"github.com/gcinema/gateway/pkg/logger"
	"go.uber.org/zap"
)

func Trace() Middleware {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			log := logger.FromContext(ctx)
			rw := response.NewResponseWriter(w)

			before := time.Now()
			log.Debug(
				">>> incoming HTTP request",
				zap.Time("time", before.UTC()))

			next.ServeHTTP(rw, r)

			log.Debug(
				">>> done HTTP request",
				zap.Int("status_code", rw.GetStatusCodeMust()),
				zap.Duration("latency", time.Since(before)))
		})
	}
}
