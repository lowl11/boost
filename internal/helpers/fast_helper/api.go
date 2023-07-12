package fast_helper

import "github.com/valyala/fasthttp"

func Write(request *fasthttp.RequestCtx, contentType string, status int, body []byte) {
	request.SetStatusCode(status)
	request.Response.Header.Set("Content-Type", contentType)
	request.SetBody(body)
}
