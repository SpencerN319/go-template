package server

import (
	"context"
	"net/http"
)

func New(ctx context.Context) http.Handler {
	mux := http.NewServeMux()
	addRoutes(mux)
	var handler http.Handler = mux
	handler = withLogging(handler)
	return handler
}
