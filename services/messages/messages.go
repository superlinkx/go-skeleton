package messages

import (
	"context"
	"fmt"

	"github.com/superlinkx/go-skeleton/models"
	"github.com/superlinkx/go-skeleton/postgres"
	"github.com/superlinkx/go-skeleton/services/messages/internal/convertors"
	"github.com/superlinkx/go-skeleton/services/messages/internal/filters"
)

type MessageStorer interface {
	GetMessageById(context.Context, int64) (postgres.Message, error)
	GetMessagesByIds(context.Context, []int64) ([]postgres.Message, error)
	GetMessageIds(context.Context) ([]int64, error)
}

type MessageService struct {
	messageRepo MessageStorer
}

func NewMessageService(repo MessageStorer) MessageService {
	return MessageService{messageRepo: repo}
}

func GetHelloMessage() models.Message {
	return models.Message{Message: "Welcome to the API"}
}

func (s MessageService) GetDatabaseMessage(ctx context.Context, id int64) (models.Message, error) {
	if message, err := s.messageRepo.GetMessageById(ctx, id); err != nil {
		return models.Message{}, fmt.Errorf("failed to get message (id=%d): %w", id, err)
	} else {
		return models.Message{Message: message.Message}, nil
	}
}

func (s MessageService) GetOddDatabaseMessages(ctx context.Context) ([]models.Message, error) {
	if messageIds, err := s.messageRepo.GetMessageIds(ctx); err != nil {
		return []models.Message{}, fmt.Errorf("failed to get message ids: %w", err)
	} else {
		oddIds := filters.FilterToOddIntegers(messageIds)
		if messages, err := s.messageRepo.GetMessagesByIds(ctx, oddIds); err != nil {
			return []models.Message{}, fmt.Errorf("failed to get messages by ids: %w", err)
		} else {
			return convertors.ConvertPostgresMessage(messages), nil
		}
	}
}
