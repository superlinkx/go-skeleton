package messages

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/superlinkx/go-skeleton/app"
	"github.com/superlinkx/go-skeleton/services/jsonservice"
	"github.com/superlinkx/go-skeleton/services/messages"
)

type EchoMessage struct {
	Message string `json:"message"`
}

type MessagesHandlers struct {
	AppContainer *app.AppContainer
}

var (
	ErrInternalServer = errors.New("internal server error")
)

func NewMessagesHandlers(appContainer *app.AppContainer) MessagesHandlers {
	return MessagesHandlers{
		AppContainer: appContainer,
	}
}

func (s MessagesHandlers) GetRoot(w http.ResponseWriter, r *http.Request) {
	var (
		greeting = messages.GetHelloMessage()
	)

	if err := jsonservice.JSONResponse(w, greeting); err != nil {
		jsonservice.JSONErrorResponse(w, 500, ErrInternalServer)
	}
}

func (s MessagesHandlers) PostEcho(w http.ResponseWriter, r *http.Request) {
	var (
		echoMessage EchoMessage
		bodyDecoder = json.NewDecoder(r.Body)
	)

	if err := bodyDecoder.Decode(&echoMessage); err != nil {
		jsonservice.JSONErrorResponse(w, 400, fmt.Errorf("invalid submission: %w", err))
	} else if err := jsonservice.JSONResponse(w, echoMessage); err != nil {
		jsonservice.JSONErrorResponse(w, 500, ErrInternalServer)
	}
}

func (s MessagesHandlers) GetDatabaseMessage(w http.ResponseWriter, r *http.Request) {
	var ()

	if message, err := messages.GetDatabaseMessage(s.AppContainer.DB); err != nil {
		jsonservice.JSONErrorResponse(w, 500, fmt.Errorf("invalid submission: %w", err))
	} else if err := jsonservice.JSONResponse(w, message); err != nil {
		jsonservice.JSONErrorResponse(w, 500, ErrInternalServer)
	}
}
