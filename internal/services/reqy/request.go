package reqy

import (
	"context"
	"net/http"
	"time"
)

type Request struct {
	baseURL string
	body    any

	headers map[string]string
	cookies map[string]string

	ctx context.Context

	client http.Client

	isXML bool

	timeout      time.Duration
	isTimeoutSet bool

	retryCount    int
	retryWaitTime time.Duration

	basicAuth *BasicAuth

	response *Response
	result   any
}

func NewRequest(baseURL string, client http.Client) *Request {
	return &Request{
		baseURL: baseURL,
		client:  client,

		headers: make(map[string]string),
		cookies: make(map[string]string),
	}
}

type BasicAuth struct {
	Username string
	Password string
}
