package server

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"

	"github.com/jennxsierra/echo-chamber/internal/logger"
)

func handleConnection(conn net.Conn) {
	clientAddr := conn.RemoteAddr().String()
	logger.LogConnection(clientAddr)

	clientLogger, clientLogFile, err := logger.CreateClientLogger(clientAddr)
	if err != nil {
		logger.LogToTerminal(fmt.Sprintf("Failed to create client logger for %s: %v", clientAddr, err))
		return
	}
	defer func() {
		logger.LogDisconnection(clientAddr)
		clientLogFile.Close()
		conn.Close()
	}()

	welcomeMessage := fmt.Sprintf(
		"Welcome to the Echo Chamber! Where input == output, and recursion is reality.\n"+
			"You are now connected as [%s]\n"+
			"Enter \"bye\" or /quit to exit.\n"+
			"Enter /help to see all commands.\n\n",
		clientAddr,
	)
	conn.Write([]byte(welcomeMessage))

	reader := bufio.NewReader(conn)
	for {
		conn.Write([]byte(fmt.Sprintf("[%s] ", clientAddr)))

		// Read input from the client
		clientMessage, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				// Client disconnected
				clientLogger.Printf("Client [%s] disconnected.", clientAddr)
				break
			}
			// Log unexpected errors and continue
			clientLogger.Printf("Error reading from client [%s]: %v", clientAddr, err)
			break
		}

		// Trim whitespace and handle the message
		clientMessage = strings.TrimSpace(clientMessage)
		clientLogger.Printf("[SENT] [%s] %s", clientAddr, clientMessage)

		// Echo the message back to the client
		serverResponse := fmt.Sprintf("[Echo Chamber] %s\n", clientMessage)
		_, writeErr := conn.Write([]byte(serverResponse))
		if writeErr != nil {
			clientLogger.Printf("Error writing to client [%s]: %v", clientAddr, writeErr)
			break
		}
		clientLogger.Printf("[RECEIVED] %s", serverResponse)

		// Handle client exit commands
		if clientMessage == "bye" || clientMessage == "/quit" {
			clientLogger.Printf("Client [%s] requested to disconnect.", clientAddr)
			break
		}
	}
}
