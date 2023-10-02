package reqy

import "net/http"

type Response struct {
	raw *http.Response

	body       []byte
	statusCode int
	status     string
}

func newResponse(raw *http.Response) *Response {
	return &Response{
		raw: raw,
	}
}

func (response *Response) Raw() *http.Response {
	return response.raw
}

func (response *Response) setBody(body []byte) *Response {
	response.body = body
	return response
}

func (response *Response) Body() []byte {
	return response.body
}

func (response *Response) setStatus(status string, code int) *Response {
	response.status = status
	response.statusCode = code
	return response
}

func (response *Response) Status() string {
	return response.status
}

func (response *Response) StatusCode() int {
	return response.statusCode
}
