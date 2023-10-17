package reqy

import (
	"context"
	"net/http"
	"time"
)

type Request struct {
	// init data
	baseURL string
	body    any
	ctx     context.Context
	client  http.Client

	// collect data
	headers map[string]string
	cookies map[string]string

	isXML bool

	timeout      time.Duration
	isTimeoutSet bool

	retryCount    int
	retryWaitTime time.Duration

	basicAuth *BasicAuth

	// cached data
	cache *http.Request

	// result data
	response      *Response
	result        any
	sendError     error
	waitForResult bool
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
