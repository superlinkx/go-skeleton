package server

import (
	"net/http"

	chi "github.com/go-chi/chi/v5"
	"github.com/superlinkx/go-skeleton/app"
	"github.com/superlinkx/go-skeleton/handlers"
	"github.com/superlinkx/go-skeleton/middleware"
)

func StartServer(appContainer app.AppContainer) {
	router := chi.NewRouter()
	middleware.Register(router)
	handlers.Register(router, &appContainer)
	http.ListenAndServe(":8080", router)
}
