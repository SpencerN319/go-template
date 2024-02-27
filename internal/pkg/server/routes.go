package server

import "net/http"

func addRoutes(mux *http.ServeMux) {
	mux.Handle("/", http.NotFoundHandler())
	mux.Handle("GET /health", handleHealth())
}
