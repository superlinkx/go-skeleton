package messages_test

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/superlinkx/go-skeleton/postgres"
	"github.com/superlinkx/go-skeleton/services/messages"
	"github.com/superlinkx/go-skeleton/services/messages/mocks"
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

func TestGetDatabaseMessage(t *testing.T) {
	var (
		mockCtrl, ctx     = gomock.WithContext(context.Background(), t)
		mockMessageGetter = mocks.NewMockMessageGetter(mockCtrl)
		messageService    = messages.NewMessageService(mockMessageGetter)
		expected          = "Hello World!"
		expectedErr       = errors.New("database pooped itself")
		returnMsg         = "Hello World!"
	)

	mockMessageGetter.EXPECT().GetMessageById(ctx, int64(0)).Return(postgres.Message{ID: 0, Message: returnMsg}, nil)

	if message, err := messageService.GetDatabaseMessage(ctx, 0); err != nil {
		t.Errorf("Unexpectedly failed to get database message: %v", err)
	} else if message.Message != expected {
		t.Errorf("Expected \"%s\", Got \"%s\"", expected, message.Message)
	}

	mockMessageGetter.EXPECT().GetMessageById(ctx, int64(2)).Return(postgres.Message{}, expectedErr)

	if _, err := messageService.GetDatabaseMessage(ctx, 2); err == nil {
		t.Errorf("Unexpectedly succeeded to get database message: %v", err)
	} else if !errors.As(err, expectedErr) {
		t.Errorf("Expected \"%s\", Got \"%s\"", expectedErr, err)
	}
}
