package logger

import (
	"log"
	"os"
	"time"
)

var logFile *os.File

// InitializeLogger sets up the log file
func InitializeLogger() error {
	var err error
	logFile, err = os.OpenFile("server.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	log.SetOutput(logFile)
	return nil
}

// CloseLogger closes the log file
func CloseLogger() {
	if logFile != nil {
		logFile.Close()
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
