package reqesty

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"github.com/lowl11/lazylog/log"
	"io"
	"net/http"
)

func (req *Request) do(ctx context.Context) error {
	var request *http.Request
	var err error

	if req.body == nil {
		request, err = http.NewRequestWithContext(ctx, req.method, req.url, nil)
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

		request, err = http.NewRequestWithContext(ctx, req.method, req.url, bytes.NewBuffer(parsedBody))
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

	req.response = response
	req.responseStatus = response.Status
	req.responseStatusCode = response.StatusCode

	responseBody, err := io.ReadAll(response.Body)
	if err != nil {
		return err
	}

	req.responseBody = responseBody

	return nil
}
