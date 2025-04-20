package logger

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

var serverLogFile *os.File

// InitializeLogger sets up the server log file
func InitializeLogger() error {
	logsDir := "logs"
	err := os.MkdirAll(logsDir, 0755)
	if err != nil {
		return fmt.Errorf("failed to create logs directory: %v", err)
	}

	serverLogPath := filepath.Join(logsDir, fmt.Sprintf("server_%s.log", time.Now().Format("20060102_150405")))
	serverLogFile, err = os.OpenFile(serverLogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("failed to create server log file: %v", err)
	}

	log.SetOutput(serverLogFile)
	log.SetFlags(0)
	log.SetPrefix("")
	return nil
}

// LogMessage logs a message with a custom timestamp format
func LogMessage(message string) {
	log.Printf("%s %s", time.Now().Format("2006/01/02 15:04:05"), message)
}

// LogToTerminal logs a message with a timestamp to the terminal
func LogToTerminal(message string) {
	timestamp := formatTimestamp()
	fmt.Printf("[%s] %s\n", timestamp, message)
}

// CloseLogger closes the server log file
func CloseLogger() {
	if serverLogFile != nil {
		serverLogFile.Close()
	}
}

// formatTimestamp returns the current timestamp in "DD-MON-YYYY HH:MM:SS" format
func formatTimestamp() string {
	return time.Now().Format("02-Jan-2006 15:04:05")
}

// LogConnection logs when a client connects
func LogConnection(clientAddr string) {
	message := fmt.Sprintf("[%s] connected to the server", clientAddr)
	LogToTerminal(message)
	LogMessage(message)
}

// LogDisconnection logs when a client disconnects
func LogDisconnection(clientAddr string) {
	message := fmt.Sprintf("[%s] disconnected from the server", clientAddr)
	LogToTerminal(message)
	LogMessage(message)
}

// CreateClientLogger creates a log file for a specific client
func CreateClientLogger(clientAddr string) (*log.Logger, *os.File, error) {
	logsDir := "logs"
	err := os.MkdirAll(logsDir, 0755)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create logs directory: %v", err)
	}

	// Sanitize clientAddr to remove periods and colons
	sanitizedClientAddr := strings.ReplaceAll(clientAddr, ".", "_")
	sanitizedClientAddr = strings.ReplaceAll(sanitizedClientAddr, ":", "_")

	clientLogPath := filepath.Join(logsDir, fmt.Sprintf("client_%s.log", sanitizedClientAddr))
	clientLogFile, err := os.OpenFile(clientLogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create client log file: %v", err)
	}

	clientLogger := log.New(clientLogFile, "", log.LstdFlags)
	return clientLogger, clientLogFile, nil
}
