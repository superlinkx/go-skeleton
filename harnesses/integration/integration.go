//go:build integration || e2e
// +build integration e2e

package integration

import (
	"database/sql"
	"embed"
	_ "embed"
	"fmt"
	"os"

	migrate "github.com/rubenv/sql-migrate"
	"github.com/superlinkx/go-skeleton/app"
	"github.com/superlinkx/go-skeleton/db"
)

func NewAppContainer() (app.AppContainer, error) {
	connStr := getIntegrationTestDatabaseURL()

	if appContainer, err := app.NewApp(connStr); err != nil {
		return app.AppContainer{}, fmt.Errorf("failed to create new app container: %w", err)
	} else {
		return appContainer, nil
	}
}

func ConnectDatabase() (*sql.DB, error) {
	connStr := getIntegrationTestDatabaseURL()

	if db, err := db.NewDatabaseConnection(connStr); err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	} else {
		return db, nil
	}
}

//go:embed migrations
var migrationFS embed.FS

func MigrateDatabase(db *sql.DB) error {
	migrations := migrate.EmbedFileSystemMigrationSource{
		FileSystem: migrationFS,
		Root:       "migrations",
	}

	if _, err := migrate.Exec(db, "postgres", migrations, migrate.Down); err != nil {
		return fmt.Errorf("failed to migrate down: %w", err)
	} else if _, err := migrate.Exec(db, "postgres", migrations, migrate.Up); err != nil {
		return fmt.Errorf("failed to migrate up: %w", err)
	} else {
		return nil
	}
}

func getIntegrationTestDatabaseURL() string {
	var connStr string

	if connStr = os.Getenv("INTEGRATION_TEST_DATABASE_URL"); connStr == "" {
		connStr = "postgres://postgres:postgres@localhost:55432/postgres?sslmode=disable"
	}

	return connStr
}
