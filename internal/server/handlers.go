package server

import (
	"fmt"
	"net"

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

	buf := make([]byte, 1024)
	for {
		conn.Write([]byte(fmt.Sprintf("[%s] ", clientAddr)))
		n, err := conn.Read(buf)
		if err != nil {
			clientLogger.Printf("Error reading from client: %v", err)
			break
		}

		clientMessage := string(buf[:n])
		clientLogger.Printf("[SENT] [%s] %s", clientAddr, clientMessage)

		// Echo the message back to the client with the server's nametag
		serverResponse := fmt.Sprintf("[Echo Chamber] %s\n", clientMessage)
		conn.Write([]byte(serverResponse))
		clientLogger.Printf("[RECEIVED] %s", serverResponse)

		if clientMessage == "bye" || clientMessage == "/quit" {
			break
		}
	}
}
