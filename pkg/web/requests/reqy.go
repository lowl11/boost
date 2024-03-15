package requests

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"github.com/lowl11/boost/log"
	"github.com/lowl11/boost/pkg/io/types"
	"io"
	"net/http"
	"strings"
	"time"
)

type Reqy struct {
	// init data
	baseURL string
	body    any
	ctx     context.Context
	client  *http.Client

	// collect data
	headers map[string]string
	cookies map[string]string

	isXML bool

	timeout      time.Duration
	isTimeoutSet bool

	retryCount    int
	retryWaitTime time.Duration

	basicAuth *basicAuth

	// cached data
	cache *http.Request

	// result data
	response      *reqyResponse
	result        any
	sendError     error
	waitForResult bool
}

func newReqy(baseURL string, client *http.Client) *Reqy {
	return &Reqy{
		baseURL: baseURL,
		client:  client,

		headers: make(map[string]string),
		cookies: make(map[string]string),
	}
}

func (req *Reqy) SetContext(ctx context.Context) *Reqy {
	req.ctx = ctx
	return req
}

func (req *Reqy) Context() context.Context {
	return req.getContext()
}

func (req *Reqy) Headers() map[string]string {
	return req.headers
}

func (req *Reqy) SetHeader(key, value string) *Reqy {
	req.headers[key] = value
	return req
}

func (req *Reqy) SetHeaders(headers map[string]string) *Reqy {
	if headers == nil {
		return req
	}

	for key, value := range headers {
		req.headers[key] = value
	}

	return req
}

func (req *Reqy) Cookies() map[string]string {
	return req.cookies
}

func (req *Reqy) SetCookie(key, value string) *Reqy {
	req.cookies[key] = value

	return req
}

func (req *Reqy) SetCookies(cookies map[string]string) *Reqy {
	if cookies == nil {
		return req
	}

	for key, value := range cookies {
		req.cookies[key] = value
	}

	return req
}

func (req *Reqy) SetBody(body any) *Reqy {
	req.body = body
	return req
}

func (req *Reqy) Body() any {
	return req.body
}

func (req *Reqy) Response() *reqyResponse {
	return req.response
}

func (req *Reqy) SetTimeout(timeout time.Duration) *Reqy {
	req.timeout = timeout
	req.isTimeoutSet = true

	return req
}

func (req *Reqy) Timeout() time.Duration {
	return req.timeout
}

func (req *Reqy) SetRetryCount(count int) *Reqy {
	req.retryCount = count
	return req
}

func (req *Reqy) SetRetryWaitTime(waitTime time.Duration) *Reqy {
	req.retryWaitTime = waitTime
	return req
}

func (req *Reqy) SetBasicAuth(basicAuth *basicAuth) *Reqy {
	req.basicAuth = basicAuth
	return req
}

func (req *Reqy) SetResult(result any) *Reqy {
	req.result = result
	req.waitForResult = true
	return req
}

func (req *Reqy) Error() error {
	return req.sendError
}

func (req *Reqy) GET(url string) (*reqyResponse, error) {
	if err := req.do(http.MethodGet, url, req.getContext()); err != nil {
		return nil, err
	}

	return req.response, nil
}

func (req *Reqy) POST(url string) (*reqyResponse, error) {
	if err := req.do(http.MethodPost, url, req.getContext()); err != nil {
		return nil, err
	}

	return req.response, nil
}

func (req *Reqy) PUT(url string) (*reqyResponse, error) {
	if err := req.do(http.MethodPut, url, req.getContext()); err != nil {
		return nil, err
	}

	return req.response, nil
}

func (req *Reqy) DELETE(url string) (*reqyResponse, error) {
	if err := req.do(http.MethodDelete, url, req.getContext()); err != nil {
		return nil, err
	}

	return req.response, nil
}

func (req *Reqy) Do(method, url string) (*reqyResponse, error) {
	if err := req.do(method, url, req.getContext()); err != nil {
		return nil, err
	}

	return req.response, nil
}

func (req *Reqy) do(method, url string, ctx context.Context) error {
	if req.retryCount > 0 { // call retry times
		if req.retryWaitTime == 0 {
			req.retryWaitTime = time.Millisecond * 100
		}

		for i := 0; i < req.retryCount; i++ {
			err := req.execute(method, url, ctx)
			if err == nil {
				return nil
			}

			req.sendError = err
			time.Sleep(req.retryWaitTime)
		}
	} else { // call once
		req.sendError = req.execute(method, url, ctx)
	}

	return req.sendError
}

func (req *Reqy) execute(method, url string, ctx context.Context) error {
	var request *http.Request
	var err error

	if req.cache == nil {
		// create new request
		request, err = req.createNewRequest(ctx, method, url)
		if err != nil {
			return err
		}

		// cache request
		req.cache = request
	} else {
		request = req.cache
	}

	// send request
	response, err := req.client.Do(request)
	if err != nil {
		return err
	}

	// save response
	req.response = newResponse(response).
		setStatus(response.Status, response.StatusCode)

	// parse response
	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}
	defer func() {
		if err = response.Body.Close(); err != nil {
			log.Warn(err, "Close request body error")
		}
	}()

	// set response bytes body
	req.response.setBody(responseBody)

	// try to unmarshal response body if response code is success
	if req.waitForResult && response.StatusCode < http.StatusBadRequest &&
		!strings.Contains(types.String(responseBody), "<!DOCTYPE html>") &&
		!strings.Contains(types.String(responseBody), "ERROR = ") {
		if err = req.unmarshal(responseBody, &req.result); err != nil {
			log.Error("Unmarshal result error:", err)
		}
	}

	return nil
}

func (req *Reqy) createNewRequest(ctx context.Context, method, url string) (*http.Request, error) {
	var request *http.Request
	var err error

	requestURL := req.baseURL + url

	if req.body == nil {
		request, err = http.NewRequestWithContext(ctx, method, requestURL, nil)
	} else {
		var parsedBody []byte

		if !req.isXML {
			parsedBody, err = json.Marshal(req.body)
			if err != nil {
				return nil, err
			}

		} else {
			parsedBody, err = xml.Marshal(req.body)
			if err != nil {
				return nil, err
			}
		}

		request, err = http.NewRequestWithContext(ctx, method, requestURL, bytes.NewBuffer(parsedBody))
	}
	if err != nil {
		return nil, err
	}

	// fill meta data
	fillHeaders(request, req.headers)
	fillCookies(request, req.cookies)

	// set basic auth
	if req.basicAuth != nil {
		request.SetBasicAuth(req.basicAuth.Username, req.basicAuth.Password)
	}

	return request, nil
}

func (req *Reqy) getContext() context.Context {
	if req.ctx == nil {
		return context.Background()
	}

	return req.ctx
}

func (req *Reqy) unmarshal(body []byte, result any) error {
	if req.isXML {
		return xml.Unmarshal(body, &result)
	}

	//if req.response.raw.Header.Get("Content-Type") != content_types.JSON {
	//req.result = type_helper.String(body, false)
	//}

	return json.Unmarshal(body, &result)
}

type basicAuth struct {
	Username string
	Password string
}

type reqyResponse struct {
	raw     *http.Response
	cookies map[string]string

	body       []byte
	statusCode int
	status     string
}

func newResponse(raw *http.Response) *reqyResponse {
	return &reqyResponse{
		raw: raw,
	}
}

func (response *reqyResponse) Raw() *http.Response {
	return response.raw
}

func (response *reqyResponse) setBody(body []byte) *reqyResponse {
	response.body = body
	return response
}

func (response *reqyResponse) Body() []byte {
	return response.body
}

func (response *reqyResponse) setStatus(status string, code int) *reqyResponse {
	response.status = status
	response.statusCode = code
	return response
}

func (response *reqyResponse) Status() string {
	return response.status
}

func (response *reqyResponse) StatusCode() int {
	return response.statusCode
}

func (response *reqyResponse) Header(key string) string {
	return response.raw.Header.Get(key)
}

func (response *reqyResponse) Headers() map[string]string {
	headers := make(map[string]string)
	for key, value := range response.raw.Header {
		headers[key] = strings.Join(value, ",")
	}
	return headers
}

func (response *reqyResponse) Cookie(key string) string {
	return response.Cookies()[key]
}

func (response *reqyResponse) Cookies() map[string]string {
	if response.cookies != nil {
		return response.cookies
	}

	for _, c := range response.raw.Cookies() {
		response.cookies[c.Name] = c.Value
	}

	return response.cookies
}

func (response *reqyResponse) ContentType() string {
	return response.Header("Content-Type")
}

func fillHeaders(request *http.Request, headers map[string]string) {
	if request == nil || headers == nil {
		return
	}

	for key, value := range headers {
		request.Header.Set(key, value)
	}
}

func fillCookies(request *http.Request, cookies map[string]string) {
	if request == nil || cookies == nil {
		return
	}

	for name, value := range cookies {
		request.AddCookie(&http.Cookie{
			Name:  name,
			Value: value,
		})
	}
}
