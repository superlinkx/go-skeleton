package messages

import (
	"context"
	"fmt"

	"github.com/superlinkx/go-skeleton/postgres"
)

type MessageRepository interface {
	GetMessage(context.Context, int64) (postgres.Message, error)
	GetMessageIds(context.Context) ([]int64, error)
}

type Message struct {
	Message string `json:"message"`
}

func GetHelloMessage() Message {
	return Message{Message: "Welcome to the API"}
}

func GetDatabaseMessage(ctx context.Context, messageRepo MessageRepository, id int64) (Message, error) {
	if message, err := messageRepo.GetMessage(ctx, id); err != nil {
		return Message{}, fmt.Errorf("failed to get message (id=%d): %w", id, err)
	} else {
		return Message{Message: message.Message}, nil
	}
}
