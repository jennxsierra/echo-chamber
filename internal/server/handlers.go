package server

import (
	"fmt"
	"net"

	"github.com/jennxsierra/echo-chamber/internal/logger"
)

func handleConnection(conn net.Conn) {
	clientAddr := conn.RemoteAddr().String()
	logger.LogConnection(clientAddr)

	// Create a client-specific logger
	clientLogger, clientLogFile, err := logger.CreateClientLogger(clientAddr)
	if err != nil {
		fmt.Printf("Failed to create client logger for %s: %v\n", clientAddr, err)
		return
	}
	defer func() {
		logger.LogDisconnection(clientAddr)
		clientLogFile.Close()
		conn.Close()
	}()

	buf := make([]byte, 1024)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			clientLogger.Printf("Error reading from client: %v", err)
			return
		}
		clientLogger.Printf("Received: %s", string(buf[:n]))

		_, err = conn.Write(buf[:n])
		if err != nil {
			clientLogger.Printf("Error writing to client: %v", err)
		}
	}
}
