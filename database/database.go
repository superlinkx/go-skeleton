package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func NewDatabaseConnection(connStr string) (*sql.DB, error) {
	if db, err := sql.Open("postgres", connStr); err != nil {
		return nil, fmt.Errorf("error opening connection to postgres: %w", err)
	} else {
		return db, nil
	}
}
