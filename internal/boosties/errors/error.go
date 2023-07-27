package errors

import "net/http"

const (
	unknownErrorType = "Unknown error type"
)

type Error struct {
	message   string
	errorType string
	httpCode  int
	context   map[string]any
}

type OutputError struct {
	Status  string         `json:"status"`
	Message string         `json:"message"`
	Type    string         `json:"type"`
	Code    int            `json:"code"`
	Context map[string]any `json:"context"`
}

func New(message string) *Error {
	return &Error{
		message:   message,
		errorType: unknownErrorType,
		httpCode:  http.StatusInternalServerError,
		context:   make(map[string]any),
	}
}
