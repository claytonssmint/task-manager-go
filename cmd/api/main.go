package main

import (
	"log"

	"github.com/claytonssmint/task-manager-go/pkg/config"
)

func main() {
	if err := config.Run(); err != nil {
		log.Fatalf("erro ao iniciar servidor: %v", err)
	}
}
