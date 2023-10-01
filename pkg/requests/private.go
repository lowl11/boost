package requests

import (
	"github.com/lowl11/boost/pkg/errors"
	"net/http"
)

func (service *Service) transport() (*http.Transport, error) {
	transport, ok := service.client.Transport.(*http.Transport)
	if !ok {
		return nil, errors.New("HTTP Client does not content *http.Transport")
	}

	return transport, nil
}
