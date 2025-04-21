package server

import "strings"

// CommandAction holds the custom response and a disconnect flag.
type CommandAction struct {
	Response        string
	DisconnectAfter bool
}

// personalityResponses maps input commands to their actions.
var personalityResponses = map[string]CommandAction{
	"hello": {Response: "[Echo Chamber] Hi there!\n\n", DisconnectAfter: false},
	"":      {Response: "[Echo Chamber] Say something...\n\n", DisconnectAfter: false},
	"bye":   {Response: "[Echo Chamber] Goodbye!\n", DisconnectAfter: true},
}

// HandlePersonalityCommand checks if a message is a custom personality command.
// It returns the response, disconnect flag, and a boolean signifying if the command was handled.
func HandlePersonalityCommand(message string) (string, bool, bool) {
	normalized := strings.TrimSpace(strings.ToLower(message))
	action, exists := personalityResponses[normalized]
	if !exists {
		return "", false, false
	}
	return action.Response, action.DisconnectAfter, true
}
