package context

import (
	"github.com/valyala/fasthttp"
	"sync"
)

type fastWriter struct {
	request    *fasthttp.RequestCtx
	wasWritten bool
	mutex      sync.Mutex
}

func newFastWriter(request *fasthttp.RequestCtx) *fastWriter {
	return &fastWriter{
		request: request,
	}
}

func (writer *fastWriter) Write(contentType string, status int, body []byte) {
	writer.mutex.Lock()
	defer writer.mutex.Unlock()

	if writer.wasWritten {
		return
	}

	writer.request.SetStatusCode(status)
	writer.request.Response.Header.Set("Content-Type", contentType)
	writer.request.SetBody(body)

	writer.wasWritten = true
}
