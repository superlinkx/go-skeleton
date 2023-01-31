package app

import (
	"github.com/superlinkx/go-skeleton/db"
	"github.com/superlinkx/go-skeleton/postgres"
)

type AppContainer struct {
	Queries *postgres.Queries
}

func SetAppDependencies(dbConnStr string) (AppContainer, error) {
	if db, err := db.NewDatabaseConnection(dbConnStr); err != nil {
		return AppContainer{}, err
	} else {
		return AppContainer{
			Queries: postgres.New(db),
		}, nil
	}
}
