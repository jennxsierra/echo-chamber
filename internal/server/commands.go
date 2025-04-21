package server

import (
	"fmt"
	"net"
	"strings"
	"time"
)

// CommandHandler defines the signature of a command function.
type CommandHandler func(args []string, conn net.Conn) (response string, disconnect bool)

// commandHandlers maps supported commands to their handlers.
var commandHandlers = map[string]CommandHandler{
	"time": func(args []string, conn net.Conn) (string, bool) {
		// Return the current server time.
		return fmt.Sprintf("[Echo Chamber] Current server time: %s\n\n", time.Now().Format(time.RFC1123)), false
	},
	"quit": func(args []string, conn net.Conn) (string, bool) {
		// Return a goodbye message and indicate that the connection should close.
		return "[Echo Chamber] Closing connection...\n\n", true
	},
	"echo": func(args []string, conn net.Conn) (string, bool) {
		// Return concatenated message arguments.
		if len(args) == 0 {
			return "[Echo Chamber] Usage: /echo message\n\n", false
		}
		return strings.Join(args, " ") + "\n\n", false
	},
	"help": func(args []string, conn net.Conn) (string, bool) {
		helpText := "[Echo Chamber] Available commands:\n" +
			"/time  - Show the current server time.\n" +
			"/quit  - Close the connection.\n" +
			"/echo  - Echo back your message. Usage: /echo message\n" +
			"/help  - Show this help message.\n\n"
		return helpText, false
	},
}

// ProcessCommand checks if the input is a command (starts with "/") and,
// if so, dispatches it to the appropriate handler. It returns the response,
// a bool flag to indicate if the connection should be disconnected, and a flag
// to indicate that the command was handled.
func ProcessCommand(input string, conn net.Conn) (string, bool, bool) {
	input = strings.TrimSpace(input)
	if !strings.HasPrefix(input, "/") {
		return "", false, false
	}
	// Remove the leading "/" before processing.
	input = input[1:]
	fields := strings.Fields(input)
	if len(fields) == 0 {
		return "Invalid command format.\n", false, true
	}
	cmdName := strings.ToLower(fields[0])
	args := fields[1:]

	handler, exists := commandHandlers[cmdName]
	if !exists {
		// Return a clear error message for unknown commands.
		return "[Echo Chamber] Unknown command. Please try again.\n\n", false, true
	}

	response, disconnect := handler(args, conn)
	return response, disconnect, true
}
