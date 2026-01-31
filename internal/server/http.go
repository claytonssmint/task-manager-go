package server

import (
	"net/http"
	"time"

	"github.com/claytonssmint/task-manager-go/internal/handlers"
)

func NewHTTPServer(addr string) *http.Server {
	mux := http.NewServeMux()

	// rotas
	mux.HandleFunc("/health", handlers.HealthHandler)

	return &http.Server{
		Addr:         addr,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
}
