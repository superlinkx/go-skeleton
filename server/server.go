package server

import (
	"net/http"
	"os"

	chi "github.com/go-chi/chi/v5"
	"github.com/superlinkx/go-skeleton/app"
	"github.com/superlinkx/go-skeleton/handlers"
	"github.com/superlinkx/go-skeleton/middleware"
)

func StartServer(appContainer app.AppContainer) {
	router := chi.NewRouter()
	middleware.Register(router)
	handlers.Register(router, &appContainer)
	if err := http.ListenAndServe(":8888", router); err != nil {
		os.Exit(1)
	}
}
