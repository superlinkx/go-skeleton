package main

import (
	"log"

	"github.com/superlinkx/go-skeleton/app"
	"github.com/superlinkx/go-skeleton/server"
)

const dbConnStr = "postgres://postgres:postgres@localhost/postgres?sslmode=disable"

func main() {
	if appContainer, err := app.NewApp(dbConnStr); err != nil {
		log.Fatalf("failed to set app dependencies: %s", err.Error())
	} else {
		server.StartServer(appContainer)
	}
}
