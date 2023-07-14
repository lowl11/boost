package reqesty

import (
	"github.com/lowl11/boost/internal/boosties/context"
	"net/http"
	"time"
)

type Request struct {
	method string
	url    string
	body   any

	headers map[string]string
	cookies map[string]string

	ctx context.Context

	client http.Client

	isXML bool

	timeout      time.Duration
	isTimeoutSet bool

	response           *http.Response
	responseBody       []byte
	responseStatus     string
	responseStatusCode int
}

func New(method, url string, client http.Client) *Request {
	return &Request{
		method: method,
		url:    url,
		client: client,

		headers: make(map[string]string),
		cookies: make(map[string]string),
	}
}
