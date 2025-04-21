package server

import (
	"strings"
	"testing"
)

func TestValidateMessage_Valid(t *testing.T) {
	validMsg := "Hello, Echo Chamber!"
	// Test with a valid message
	validated, err := ValidateMessage(validMsg)
	if err != nil {
		t.Errorf("Expected valid message, got error: %v", err)
	}
	if validated != validMsg {
		t.Errorf("Expected %q but got %q", validMsg, validated)
	}
}

func TestValidateMessage_TooLong(t *testing.T) {
	// Create a message that exceeds MaxMessageSize by one byte.
	longMessage := strings.Repeat("a", MaxMessageSize+1)
	_, err := ValidateMessage(longMessage)
	if err == nil {
		t.Errorf("Expected error for message longer than %d bytes, got nil", MaxMessageSize)
	}
}
