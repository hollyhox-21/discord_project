package http

import (
	"fmt"
	"net/http"
	"runtime/debug"

	"github.com/hollyhox-21/discord_project/libraries/logger"
)

// Recover middleware
func Recover(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if p := recover(); p != nil {
				logger.Errorw(
					r.Context(), fmt.Sprintf("recovered from panic: %v", p),
					"stack_trace", string(debug.Stack()),
					"panic", true,
					"component", "http_recover_middleware",
				)

				w.WriteHeader(http.StatusInternalServerError)
				_, _ = w.Write([]byte("recover: unexpected server error"))

			}
		}()
		next.ServeHTTP(w, r)
	})
}
