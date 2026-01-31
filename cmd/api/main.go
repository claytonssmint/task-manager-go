package main

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

func main() {
	svr := server.New()

	// Rodar o servidor em goroutine
	go func() {
		log.Println("Servidor iniciado na porta 8080")
		if err := svr.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("erro ao iniciar servidor: %v", err)
		}
	}()

	//Canal para capturar sinais do sistema
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// Bloqueia até receber sinal
	<-stop
	log.Println("Recebido sinal de Shutdown, encerrando servidor...")

	// Contexto com timeout para finalizar requests
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := svr.Shutdown(ctx); err != nil {
		log.Printf("erro no shutdown: %v", err)
	} else {
		log.Println("Servidor finalizado com sucesso")
	}
}
