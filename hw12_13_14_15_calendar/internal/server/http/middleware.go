package http

import (
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/rs/zerolog/log"
)

type responseWriterDecorator struct {
	http.ResponseWriter
	statusCode int
}

func (w *responseWriterDecorator) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
	w.statusCode = statusCode
}

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()

		decoratedWriter := wrapResponseWriter(w)
		next.ServeHTTP(decoratedWriter, r)

		msg := strings.Join([]string{
			r.RemoteAddr,
			start.String(),
			r.Method,
			r.URL.Path,
			r.Proto,
			strconv.Itoa(decoratedWriter.statusCode),
			time.Since(start).String(),
			r.UserAgent(),
		}, " ")

		log.Info().Msg(msg)
	})
}

func wrapResponseWriter(w http.ResponseWriter) *responseWriterDecorator {
	return &responseWriterDecorator{
		ResponseWriter: w,
		statusCode:     http.StatusOK,
	}
}
