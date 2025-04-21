package server

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"strings"
	"time"

	"github.com/jennxsierra/echo-chamber/internal/logger"
)

// readClientMessage wraps reading a message and sets a read deadline.
func readClientMessage(conn net.Conn, reader *bufio.Reader, timeout time.Duration) (string, error) {
	// Set the read deadline to handle inactivity.
	conn.SetReadDeadline(time.Now().Add(timeout))
	message, err := reader.ReadString('\n')
	if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
		return "", fmt.Errorf("inactivity timeout: %w", err)
	}
	return message, err
}

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
	const inactivityTimeout = 30 * time.Second
	for {
		conn.Write([]byte(fmt.Sprintf("[%s] ", clientAddr)))
		clientMessage, err := readClientMessage(conn, reader, inactivityTimeout)
		if err != nil {
			if strings.Contains(err.Error(), "inactivity timeout") {
				timeoutMsg := fmt.Sprintf("Client [%s] disconnected due to inactivity.\n", clientAddr)
				conn.Write([]byte(timeoutMsg))
				if tcpConn, ok := conn.(*net.TCPConn); ok {
					tcpConn.SetLinger(0)
				}
				clientLogger.Print(timeoutMsg)
			} else if err == io.EOF || strings.Contains(err.Error(), "closed") {
				// Handle client disconnection gracefully.
				clientLogger.Printf("Client [%s] disconnected.", clientAddr)
			} else {
				clientLogger.Printf("Error reading from client [%s]: %v", clientAddr, err)
			}
			break
		}

		clientMessage = strings.TrimSpace(clientMessage)
		clientLogger.Printf("[SENT] [%s] %s", clientAddr, clientMessage)

		// Handle client exit commands.
		if clientMessage == "bye" || clientMessage == "/quit" {
			goodbyeMessage := "[Echo Chamber] Goodbye!\n"
			clientLogger.Printf("[RECEIVED] %s", goodbyeMessage)
			clientLogger.Printf("Client [%s] has exited the Echo Chamber.", clientAddr)

			// Send the goodbye message.
			_, err := conn.Write([]byte(goodbyeMessage))
			if err != nil {
				clientLogger.Printf("Error sending goodbye message to [%s]: %v", clientAddr, err)
			}

			if tcpConn, ok := conn.(*net.TCPConn); ok {
				tcpConn.CloseWrite()
				// Allow time for the goodbye message to be transmitted.
				time.Sleep(100 * time.Millisecond)
				// Now force an immediate disconnect to end netcat without extra input.
				tcpConn.SetLinger(0)
			}
			break
		}

		// Echo the message back to the client.
		serverResponse := fmt.Sprintf("[Echo Chamber] %s\n", clientMessage)
		_, writeErr := conn.Write([]byte(serverResponse))
		if writeErr != nil {
			clientLogger.Printf("Error writing to client [%s]: %v", clientAddr, writeErr)
			break
		}
		clientLogger.Printf("[RECEIVED] %s", serverResponse)
	}
}
