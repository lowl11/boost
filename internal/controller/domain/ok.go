package domain

type JustOK struct {
	Message string `json:"message"`
}

func NewJustOK() JustOK {
	return JustOK{
		Message: "OK",
	}
}

type WrappedOK struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
	Body    any    `json:"body"`
}

func NewWrappedOK(body any, message ...string) WrappedOK {
	var messageText string

	if len(message) > 0 {
		messageText = message[0]
	}

	return WrappedOK{
		Status:  "OK",
		Message: messageText,
		Body:    body,
	}
}
