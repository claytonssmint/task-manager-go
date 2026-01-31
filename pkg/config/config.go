package config

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/claytonssmint/task-manager-go/internal/server"
)

func Run() error {

	httpServer := server.NewHTTPServer(":8080")

	// Canal para capturar sinais do SO
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Sobe o servidor em goroutine
	go func() {
		log.Println("🚀 Servidor rodando na porta 8080")
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("erro ao subir servidor: %v", err)
		}
	}()

	// Fica bloqueado esperando sinal
	<-stop
	log.Println("🛑 Encerrando servidor...")

	// Contexto com timeout para finalizar requisições
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := httpServer.Shutdown(ctx); err != nil {
		log.Fatalf("erro no shutdown: %v", err)
	}

	log.Println("✅ Servidor encerrado com sucesso")
	return nil

}
