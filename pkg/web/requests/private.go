package requests

import (
	"net/http"
)

func (service *Service) transport() *http.Transport {
	transport, ok := service.client.Transport.(*http.Transport)
	if !ok {
		transport = &http.Transport{}
	}

	return transport
}
