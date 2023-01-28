package messages_test

import (
	"testing"

	"github.com/superlinkx/go-skeleton/services/messages"
)

func TestGetHelloMessage(t *testing.T) {
	var (
		expected = messages.Message{Message: "Welcome to the API"}
		result   = messages.GetHelloMessage()
	)

	if expected != result {
		t.Errorf("got %v, want %v", result, expected)
		t.Fail()
	}
}
