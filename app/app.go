package app

import (
	"database/sql"

	"github.com/superlinkx/go-skeleton/database"
)

type AppContainer struct {
	DB *sql.DB
}

func SetAppDependencies(dbConnStr string) (AppContainer, error) {
	if db, err := database.NewDatabaseConnection(dbConnStr); err != nil {
		return AppContainer{}, err
	} else {
		return AppContainer{
			DB: db,
		}, nil
	}
}
