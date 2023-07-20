//go:build e2e
// +build e2e

package messages_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"github.com/superlinkx/go-skeleton/handlers/messages"
	"github.com/superlinkx/go-skeleton/harnesses/integration"
	"github.com/superlinkx/go-skeleton/models"
)

type urlParams struct {
	Key   string
	Value string
}

func TestMessagesHandlers_GetRoot(t *testing.T) {
	appContainer, err := integration.NewAppContainer()
	require.NoError(t, err)

	messagesHandlers := messages.NewMessagesHandlers(&appContainer)

	req, err := http.NewRequest("GET", "/", nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(messagesHandlers.GetRoot)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.Contains(t, rr.Body.String(), "Welcome to the API")
}

func TestMessagesHandlers_PostEcho(t *testing.T) {
	var (
		jsonMessage = "{\"message\": \"Hello, World!\"}"
	)

	appContainer, err := integration.NewAppContainer()
	require.NoError(t, err)

	messagesHandlers := messages.NewMessagesHandlers(&appContainer)

	req, err := http.NewRequest("POST", "/echo", strings.NewReader(jsonMessage))
	require.NoError(t, err)

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(messagesHandlers.PostEcho)
	handler.ServeHTTP(rr, req)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.JSONEq(t, jsonMessage, rr.Body.String())
}

func TestMessagesHandlers_GetDatabaseMessage(t *testing.T) {
	var (
		expectedStruct = models.Message{Message: "Hello Database 1!"}
	)

	appContainer, err := integration.NewAppContainer()
	require.NoError(t, err)

	messagesHandlers := messages.NewMessagesHandlers(&appContainer)

	req, err := http.NewRequest("GET", "/database-message/{messageId}", nil)
	require.NoError(t, err)

	rr := httptest.NewRecorder()

	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("messageId", "1")

	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	handler := http.HandlerFunc(messagesHandlers.GetDatabaseMessage)
	handler.ServeHTTP(rr, req)

	bytes, err := json.Marshal(expectedStruct)
	require.NoError(t, err)

	assert.Equal(t, http.StatusOK, rr.Code)
	assert.JSONEq(t, string(bytes), rr.Body.String())
}
