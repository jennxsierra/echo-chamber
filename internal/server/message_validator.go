package server

import "fmt"

// MaxMessageSize defines the maximum allowed message length in bytes.
const MaxMessageSize = 1024

// ValidateMessage returns an error if the message exceeds MaxMessageSize.
func ValidateMessage(message string) (string, error) {
	if len(message) > MaxMessageSize {
		return "", fmt.Errorf("message length %d exceeds maximum allowed %d bytes", len(message), MaxMessageSize)
	}
	return message, nil
}
