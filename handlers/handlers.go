//go:generate mockgen -destination=mocks/mock_routerhandler.go -package=mocks github.com/superlinkx/go-skeleton/handlers RouterHandler
package handlers

import (
	"net/http"

	"github.com/superlinkx/go-skeleton/app"
	"github.com/superlinkx/go-skeleton/handlers/messages"
)

type RouterHandler interface {
	Get(string, http.HandlerFunc)
	Post(string, http.HandlerFunc)
}

func Register(router RouterHandler, appContainer *app.AppContainer) {
	messagesHandlers := messages.NewMessagesHandlers(appContainer)
	router.Get("/", messagesHandlers.GetRoot)
	router.Post("/echo", messagesHandlers.PostEcho)
	router.Get("/database-message", messagesHandlers.GetDatabaseMessage)
}
