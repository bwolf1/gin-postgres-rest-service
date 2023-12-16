package main

import (
	"os"

	"log"

	"github.com/bwolf1/gin-postgres-rest-service/pkg/config"
	"github.com/bwolf1/gin-postgres-rest-service/pkg/server"
)

func main() {
	// Load the config
	cfg, err := config.New()
	if err != nil {
		log.Printf("failed to load config: %v", err)
		os.Exit(1)
	}

	// Start the API server
	if _, err := server.New(cfg); err != nil {
		log.Printf("failed to start server: %v", err)
		os.Exit(1)
	}
}
