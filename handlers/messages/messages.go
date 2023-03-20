package messages

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/superlinkx/go-skeleton/app"
	"github.com/superlinkx/go-skeleton/handlers/utils"
	"github.com/superlinkx/go-skeleton/services/messages"
)

type EchoMessage struct {
	Message string `json:"message"`
}

type MessagesHandlers struct {
	App *app.AppContainer
}

var (
	ErrInternalServer = errors.New("internal server error")
)

func NewMessagesHandlers(appContainer *app.AppContainer) MessagesHandlers {
	return MessagesHandlers{
		App: appContainer,
	}
}

func (s MessagesHandlers) GetRoot(w http.ResponseWriter, r *http.Request) {
	var (
		greeting = messages.GetHelloMessage()
	)

	if err := utils.JSONResponse(w, greeting); err != nil {
		utils.JSONErrorResponse(w, 500, ErrInternalServer)
	}
}

func (s MessagesHandlers) PostEcho(w http.ResponseWriter, r *http.Request) {
	var (
		echoMessage EchoMessage
		bodyDecoder = json.NewDecoder(r.Body)
	)

	if err := bodyDecoder.Decode(&echoMessage); err != nil {
		utils.JSONErrorResponse(w, 400, fmt.Errorf("invalid submission: %w", err))
	} else if err := utils.JSONResponse(w, echoMessage); err != nil {
		utils.JSONErrorResponse(w, 500, ErrInternalServer)
	}
}

func (s MessagesHandlers) GetDatabaseMessage(w http.ResponseWriter, r *http.Request) {
	if messageId, err := strconv.Atoi(chi.URLParam(r, "messageId")); err != nil {
		utils.JSONErrorResponse(w, 400, fmt.Errorf("non-integer id specified: %w", err))
	} else if message, err := s.App.GetDatabaseMessage(r.Context(), int64(messageId)); err != nil {
		utils.JSONErrorResponse(w, 500, fmt.Errorf("invalid submission: %w", err))
	} else if err := utils.JSONResponse(w, message); err != nil {
		utils.JSONErrorResponse(w, 500, ErrInternalServer)
	}
}

func (s MessagesHandlers) GetOddDatabaseMessages(w http.ResponseWriter, r *http.Request) {
	if message, err := s.App.GetOddDatabaseMessages(r.Context()); err != nil {
		utils.JSONErrorResponse(w, 500, fmt.Errorf("invalid submission: %w", err))
	} else if err := utils.JSONResponse(w, message); err != nil {
		utils.JSONErrorResponse(w, 500, ErrInternalServer)
	}
}
