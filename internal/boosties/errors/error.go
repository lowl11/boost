package errors

import "net/http"

const (
	unknownErrorType = "Unknown error type"
)

type Error struct {
	message   string
	errorType string
	httpCode  int
}

func New(message string) *Error {
	return &Error{
		message:   message,
		errorType: unknownErrorType,
		httpCode:  http.StatusInternalServerError,
	}
}
