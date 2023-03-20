//go:generate mockgen -destination=mocks/mock_messagegetter.go -package=mocks github.com/superlinkx/go-skeleton/services/messages MessageGetter
package messages

import (
	"context"
	"fmt"

	"github.com/superlinkx/go-skeleton/postgres"
)

type MessageGetter interface {
	GetMessageById(context.Context, int64) (postgres.Message, error)
	GetMessagesByIds(context.Context, []int64) ([]postgres.Message, error)
	GetMessageIds(context.Context) ([]int64, error)
}

type MessageService struct {
	messageRepo MessageGetter
}

type Message struct {
	Message string `json:"message"`
}

func NewMessageService(repo MessageGetter) MessageService {
	return MessageService{messageRepo: repo}
}

func GetHelloMessage() Message {
	return Message{Message: "Welcome to the API"}
}

func (s MessageService) GetDatabaseMessage(ctx context.Context, id int64) (Message, error) {
	if message, err := s.messageRepo.GetMessageById(ctx, id); err != nil {
		return Message{}, fmt.Errorf("failed to get message (id=%d): %w", id, err)
	} else {
		return Message{Message: message.Message}, nil
	}
}

func (s MessageService) GetOddDatabaseMessages(ctx context.Context) ([]Message, error) {
	if messageIds, err := s.messageRepo.GetMessageIds(ctx); err != nil {
		return []Message{}, fmt.Errorf("failed to get message ids: %w", err)
	} else {
		oddIds := filterOddMessageIds(messageIds)
		if messages, err := s.messageRepo.GetMessagesByIds(ctx, oddIds); err != nil {
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
