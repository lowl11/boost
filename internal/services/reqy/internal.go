package reqy

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"github.com/lowl11/lazylog/log"
	"io"
	"net/http"
)

func (req *Request) do(method, url string, ctx context.Context) error {
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
				return err
			}

		} else {
			parsedBody, err = xml.Marshal(req.body)
			if err != nil {
				return err
			}
		}

		request, err = http.NewRequestWithContext(ctx, method, requestURL, bytes.NewBuffer(parsedBody))
	}
	if err != nil {
		return err
	}

	response, err := req.client.Do(request)
	if err != nil {
		return err
	}
	defer func() {
		if err = response.Body.Close(); err != nil {
			log.Warn(err, "Close request body error")
		}
	}()

	req.response = newResponse(response).
		setStatus(response.Status, response.StatusCode)

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	req.response.setBody(responseBody)

	if err = req.unmarshal(responseBody, &req.result); err != nil {
		return err
	}

	return nil
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

	return json.Unmarshal(body, &result)
}
