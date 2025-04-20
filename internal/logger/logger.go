package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"
)

var serverLogFile *os.File

// InitializeLogger sets up the server log file
func InitializeLogger() error {
	// Create logs directory if it doesn't exist
	logsDir := "logs"
	err := os.MkdirAll(logsDir, 0755)
	if err != nil {
		return fmt.Errorf("failed to create logs directory: %v", err)
	}

	// Create server log file
	serverLogPath := filepath.Join(logsDir, fmt.Sprintf("server_%s.log", time.Now().Format("20060102_150405")))
	serverLogFile, err = os.OpenFile(serverLogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to create server log file: %v", err)
	}
	log.SetOutput(serverLogFile)
	return nil
}

// CloseLogger closes the server log file
func CloseLogger() {
	if serverLogFile != nil {
		serverLogFile.Close()
	}
}

// LogConnection logs when a client connects
func LogConnection(clientAddr string) {
	log.Printf("[%s] Client connected: %s\n", time.Now().Format(time.RFC3339), clientAddr)
}

// LogDisconnection logs when a client disconnects
func LogDisconnection(clientAddr string) {
	log.Printf("[%s] Client disconnected: %s\n", time.Now().Format(time.RFC3339), clientAddr)
}

// CreateClientLogger creates a log file for a specific client
func CreateClientLogger(clientAddr string) (*log.Logger, *os.File, error) {
	// Create logs directory if it doesn't exist
	logsDir := "logs"
	err := os.MkdirAll(logsDir, 0755)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create logs directory: %v", err)
	}

	// Create client log file
	clientLogPath := filepath.Join(logsDir, fmt.Sprintf("client_%s.log", clientAddr))
	clientLogFile, err := os.OpenFile(clientLogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create client log file: %v", err)
	}

	// Create a logger for the client
	clientLogger := log.New(clientLogFile, "", log.LstdFlags)
	return clientLogger, clientLogFile, nil
}
