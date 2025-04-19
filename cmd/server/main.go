package main

import (
	"log"

	"github.com/jennxsierra/echo-chamber/internal/config"
	"github.com/jennxsierra/echo-chamber/internal/logger"
	"github.com/jennxsierra/echo-chamber/internal/server"
)

func main() {
	// Initialize configuration
	cfg := config.LoadConfig()

	// Initialize logger
	err := logger.InitializeLogger()
	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}
	defer logger.CloseLogger()

	// Start the server
	server.Start(cfg.Port)
}
