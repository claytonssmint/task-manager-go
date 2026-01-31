package server

import (
	"log"
	"net/http"
)

func New() *http.Server {
	mux := http.NewServeMux()

	//health será registrado depois
	mux.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	server := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	log.Println("HTTP server configurado na porta 8080")

	return server
}
