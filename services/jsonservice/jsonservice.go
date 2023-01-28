package jsonservice

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var (
	DefaultInternalServerError = []byte(`{"error_message": "internal server error", "code": 500}`)
)

type ErrorMessage struct {
	ErrorMessage string `json:"error_message"`
	StatusCode   int    `json:"status_code"`
}

func JSONResponse(w http.ResponseWriter, message any) error {
	if result, err := json.Marshal(message); err != nil {
		return fmt.Errorf("error marshalling JSON response: %w", err)
	} else {
		w.Write(result)
		return nil
	}
}

func JSONErrorResponse(w http.ResponseWriter, code int, err error) {
	var (
		errorMessage = ErrorMessage{
			ErrorMessage: err.Error(),
			StatusCode:   code,
		}
	)

	if result, err := json.Marshal(errorMessage); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(DefaultInternalServerError)
	} else {
		w.WriteHeader(code)
		w.Write(result)
	}
}
