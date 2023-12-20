package requests

import (
	"github.com/lowl11/boost/internal/services/web/reqy"
	"net/http"
	"time"
)

type Service struct {
	baseURL string

	headers map[string]string
	cookies map[string]string

	timeout *time.Duration

	retryCount    int
	retryWaitTime time.Duration

	basicAuth *reqy.BasicAuth

	client *http.Client
}

func New() *Service {
	return &Service{
		headers: make(map[string]string),
		cookies: make(map[string]string),

		retryWaitTime: time.Millisecond * 100,

		client: &http.Client{},
	}
}
