package server

import (
	"bytes"
	"io"
	"log/slog"
	"net/http"
	"time"
)

const (
	ContextTraceId string = "trace-id"
)

func withLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/favicon.ico", "/health":
			next.ServeHTTP(w, r)
		default:
			start := time.Now().UTC()
			var body []byte
			r.Body = io.NopCloser(io.TeeReader(r.Body, bytes.NewBuffer(body)))
			slog.Debug(
				r.URL.String(),
				slog.String("path", r.URL.EscapedPath()),
				slog.String("method", r.Method),
				slog.String("body", string(body)),
				slog.Any("form", r.Form),
				slog.Duration("duration ns", time.Duration(time.Since(start).Nanoseconds())),
			)
			next.ServeHTTP(w, r)
		}
	})
}
