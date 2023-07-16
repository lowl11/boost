package reqy

import (
	"context"
	"net/http"
	"time"
)

func (req *Request) SetContext(ctx context.Context) *Request {
	req.ctx = ctx
	return req
}

func (req *Request) Context() context.Context {
	return req.getContext()
}

func (req *Request) Headers() map[string]string {
	return req.headers
}

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

func (req *Request) Cookies() map[string]string {
	return req.cookies
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

func (req *Request) Response() *Response {
	return req.response
}

func (req *Request) SetTimeout(timeout time.Duration) *Request {
	req.timeout = timeout
	req.isTimeoutSet = true

	return req
}

func (req *Request) Timeout() time.Duration {
	return req.timeout
}

func (req *Request) SetRetryCount(count int) *Request {
	req.retryCount = count
	return req
}

func (req *Request) SetRetryWaitTime(waitTime time.Duration) *Request {
	req.retryWaitTime = waitTime
	return req
}

func (req *Request) SetResult(result any) *Request {
	req.result = result
	return req
}

func (req *Request) GET(url string) (*Response, error) {
	if err := req.do(http.MethodGet, url, req.getContext()); err != nil {
		return nil, err
	}

	return req.response, nil
}

func (req *Request) POST(url string) (*Response, error) {
	if err := req.do(http.MethodPost, url, req.getContext()); err != nil {
		return nil, err
	}

	return req.response, nil
}

func (req *Request) PUT(url string) (*Response, error) {
	if err := req.do(http.MethodPut, url, req.getContext()); err != nil {
		return nil, err
	}

	return req.response, nil
}

func (req *Request) DELETE(url string) (*Response, error) {
	if err := req.do(http.MethodDelete, url, req.getContext()); err != nil {
		return nil, err
	}

	return req.response, nil
}
