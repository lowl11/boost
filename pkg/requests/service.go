package requests

import (
	"net/http"
	"time"
)

type Service struct {
	baseURL string

	headers map[string]string
	cookies map[string]string

	timeout *time.Duration

	client http.Client
}

func New() *Service {
	return &Service{
		headers: make(map[string]string),
		cookies: make(map[string]string),

		client: http.Client{},
	}
}
