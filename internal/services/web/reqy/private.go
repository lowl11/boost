package reqy

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"github.com/lowl11/boost/data/enums/content_types"
	"github.com/lowl11/boost/internal/helpers/request_helper"
	"github.com/lowl11/boost/internal/helpers/type_helper"
	"github.com/lowl11/boost/log"
	"io"
	"net/http"
	"strings"
	"time"
)

func (req *Request) do(method, url string, ctx context.Context) error {
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

func (req *Request) execute(method, url string, ctx context.Context) error {
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
		!strings.Contains(type_helper.ToString(responseBody, false), "<!DOCTYPE html>") &&
		!strings.Contains(type_helper.ToString(responseBody, false), "ERROR = ") {
		if err = req.unmarshal(responseBody, &req.result); err != nil {
			log.Error(err, "Unmarshal result error")
		}
	}

	return nil
}

func (req *Request) createNewRequest(ctx context.Context, method, url string) (*http.Request, error) {
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
	request_helper.FillHeaders(request, req.headers)
	request_helper.FillCookies(request, req.cookies)

	// set basic auth
	if req.basicAuth != nil {
		request.SetBasicAuth(req.basicAuth.Username, req.basicAuth.Password)
	}

	return request, nil
}

func (req *Request) getContext() context.Context {
	if req.ctx == nil {
		return context.Background()
	}

	return req.ctx
}

func (req *Request) unmarshal(body []byte, result any) error {
	if req.isXML {
		return xml.Unmarshal(body, &result)
	}

	if req.response.raw.Header.Get("Content-Type") != content_types.JSON {
		req.result = type_helper.ToString(body, false)
	}

	return json.Unmarshal(body, &result)
}
