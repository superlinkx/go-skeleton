package main

import (
	"log"
	"os"

	"github.com/superlinkx/go-skeleton/app"
	"github.com/superlinkx/go-skeleton/server"
)

func main() {
	var dbConnStr string

	if dbConnStr = os.Getenv("DB_URL"); dbConnStr == "" {
		dbConnStr = "postgres://postgres:postgres@localhost/postgres?sslmode=disable"
	}

	if appContainer, err := app.NewApp(dbConnStr); err != nil {
		log.Fatalf("failed to set app dependencies: %s", err.Error())
	} else {
		server.StartServer(appContainer)
	}
}
