package messages

import (
	"context"
	"fmt"

	"github.com/superlinkx/go-skeleton/postgres"
)

type MessageRepository interface {
	GetMessageById(context.Context, int64) (postgres.Message, error)
	GetMessagesByIds(context.Context, []int64) ([]postgres.Message, error)
	GetMessageIds(context.Context) ([]int64, error)
}

type Message struct {
	Message string `json:"message"`
}

func GetHelloMessage() Message {
	return Message{Message: "Welcome to the API"}
}

func GetDatabaseMessage(ctx context.Context, messageRepo MessageRepository, id int64) (Message, error) {
	if message, err := messageRepo.GetMessageById(ctx, id); err != nil {
		return Message{}, fmt.Errorf("failed to get message (id=%d): %w", id, err)
	} else {
		return Message{Message: message.Message}, nil
	}
}

func GetOddDatabaseMessages(ctx context.Context, messageRepo MessageRepository) ([]Message, error) {
	if messageIds, err := messageRepo.GetMessageIds(ctx); err != nil {
		return []Message{}, fmt.Errorf("failed to get message ids: %w", err)
	} else {
		oddIds := filterOddMessageIds(messageIds)
		if messages, err := messageRepo.GetMessagesByIds(ctx, oddIds); err != nil {
			return []Message{}, fmt.Errorf("failed to get messages by ids: %w", err)
		} else {
			resultMessages := make([]Message, 0, len(messages))
			for _, message := range messages {
				resultMessages = append(resultMessages, Message{Message: message.Message})
			}
			return resultMessages, nil
		}
	}
}

func filterOddMessageIds(messageIds []int64) []int64 {
	var (
		oddIds = make([]int64, 0, len(messageIds)/2)
	)

	for _, id := range messageIds {
		if id%2 == 1 {
			oddIds = append(oddIds, id)
		}
	}

	return oddIds
}
