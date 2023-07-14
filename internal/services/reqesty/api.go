package reqesty

import (
	"context"
	"net/http"
	"time"
)

func (req *Request) SetHeader(key, value string) *Request {
	req.headers[key] = value
	return req
}

func (req *Request) SetHeaders(headers map[string]string) *Request {
	if headers == nil {
		return req
	}

	for key, value := range headers {
		req.headers[key] = value
	}

	return req
}

func (req *Request) SetCookie(key, value string) *Request {
	req.cookies[key] = value

	return req
}

func (req *Request) SetCookies(cookies map[string]string) *Request {
	if cookies == nil {
		return req
	}

	for key, value := range cookies {
		req.cookies[key] = value
	}

	return req
}

func (req *Request) SetBody(body any) *Request {
	req.body = body
	return req
}

func (req *Request) Body() any {
	return req.body
}

func (req *Request) Response() *http.Response {
	return req.response
}

func (req *Request) ResponseBody() []byte {
	return req.responseBody
}

func (req *Request) StatusCode() int {
	return req.responseStatusCode
}

func (req *Request) Status() string {
	return req.responseStatus
}

func (req *Request) Timeout(timeout time.Duration) *Request {
	req.timeout = timeout
	req.isTimeoutSet = true

	return req
}

func (req *Request) Do() error {
	return req.do(context.Background())
}

func (req *Request) DoCtx(ctx context.Context) error {
	return req.do(ctx)
}
