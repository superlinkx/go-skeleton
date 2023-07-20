package app

import (
	"context"

	"github.com/superlinkx/go-skeleton/db"
	"github.com/superlinkx/go-skeleton/models"
	"github.com/superlinkx/go-skeleton/postgres"
	"github.com/superlinkx/go-skeleton/services/messages"
)

type AppContainer struct {
	messages MessageServicer
}

type MessageServicer interface {
	GetDatabaseMessage(context.Context, int64) (models.Message, error)
	GetOddDatabaseMessages(context.Context) ([]models.Message, error)
}

func NewApp(dbConnStr string) (AppContainer, error) {
	if db, err := db.NewDatabaseConnection(dbConnStr); err != nil {
		return AppContainer{}, err
	} else {
		queries := postgres.New(db)
		return AppContainer{
			messages: messages.NewMessageService(queries),
		}, nil
	}
}

func (s AppContainer) GetDatabaseMessage(ctx context.Context, id int64) (models.Message, error) {
	return s.messages.GetDatabaseMessage(ctx, id)
}

func (s AppContainer) GetOddDatabaseMessages(ctx context.Context) ([]models.Message, error) {
	return s.messages.GetOddDatabaseMessages(ctx)
}
