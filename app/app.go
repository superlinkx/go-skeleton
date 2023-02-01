package app

import (
	"context"

	"github.com/superlinkx/go-skeleton/db"
	"github.com/superlinkx/go-skeleton/postgres"
	"github.com/superlinkx/go-skeleton/services/messages"
)

type AppContainer struct {
	queries *postgres.Queries
}

func NewApp(dbConnStr string) (AppContainer, error) {
	if db, err := db.NewDatabaseConnection(dbConnStr); err != nil {
		return AppContainer{}, err
	} else {
		return AppContainer{
			queries: postgres.New(db),
		}, nil
	}
}

func (s AppContainer) GetDatabaseMessage(ctx context.Context, id int64) (messages.Message, error) {
	return messages.GetDatabaseMessage(ctx, s.queries, id)
}

func (s AppContainer) GetOddDatabaseMessages(ctx context.Context) ([]messages.Message, error) {
	return messages.GetOddDatabaseMessages(ctx, s.queries)
}
