package fast_writer

import (
	"github.com/valyala/fasthttp"
	"sync"
)

type Writer struct {
	request    *fasthttp.RequestCtx
	wasWritten bool
	mutex      sync.Mutex
}

func New(request *fasthttp.RequestCtx) *Writer {
	return &Writer{
		request: request,
	}
}
