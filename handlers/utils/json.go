package utils

import (
	"encoding/json"
	"fmt"
	"log"
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
		return fmt.Errorf("failed to marshal JSON response: %w", err)
	} else if _, err := w.Write(result); err != nil {
		return fmt.Errorf("failed to write JSON result: %w", err)
	} else {
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
		if _, err := w.Write(DefaultInternalServerError); err != nil {
			log.Printf("Failed to write JSON result while processing an error writing an: %s", err)
		}
	} else {
		w.WriteHeader(code)
		if _, err := w.Write(result); err != nil {
			log.Printf("Failed to write error JSON result: %s", err)
		}
	}
}
