package context

import "github.com/lowl11/boost/pkg/system/types"

type createdWithIDNumber struct {
	ID int `json:"id"`
}

type createdWithIDString struct {
	ID string `json:"id"`
}

func newCreatedWithID(id any) any {
	numberID, ok := id.(int)
	if ok {
		return createdWithIDNumber{
			ID: numberID,
		}
	}

	stringID, ok := id.(string)
	if ok {
		return createdWithIDString{
			ID: stringID,
		}
	}

	return createdWithIDString{
		ID: types.ToString(id),
	}
}

type notFoundMessage struct {
	Message string `json:"message"`
}

func newNotFoundMessage(message string) notFoundMessage {
	return notFoundMessage{
		Message: message,
	}
}

type justOK struct {
	Message string `json:"message"`
}

func newJustOK() justOK {
	return justOK{
		Message: "OK",
	}
}

type wrappedOK struct {
	Status  string `json:"status"`
	Message string `json:"message,omitempty"`
	Body    any    `json:"body"`
}

func newWrappedOK(body any, message ...string) wrappedOK {
	var messageText string

	if len(message) > 0 {
		messageText = message[0]
	}

	return wrappedOK{
		Status:  "OK",
		Message: messageText,
		Body:    body,
	}
}
