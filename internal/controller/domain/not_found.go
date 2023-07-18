package domain

type NotFoundMessage struct {
	Message string `json:"message"`
}

func NewNotFoundMessage(message string) NotFoundMessage {
	return NotFoundMessage{
		Message: message,
	}
}
