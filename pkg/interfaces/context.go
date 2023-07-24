package interfaces

import (
	"github.com/valyala/fasthttp"
)

// Context is interface of context which comes to Boost Handler
type Context interface {
	// Request returns Fast HTTP request object
	Request() *fasthttp.Request
	// Response returns Fast HTTP response object
	Response() *fasthttp.Response

	// Method returns request method
	Method() string

	// Param returns variable param from query
	Param(name string) Param
	// Params returns all variable params
	Params() map[string]Param
	// QueryParam returns query param from query
	QueryParam(name string) Param
	// QueryParams returns all query params
	QueryParams() map[string]Param
	// Header returns header value from given name
	Header(name string) string
	// Headers returns all map of headers
	Headers() map[string]string
	// Cookie returns cookie value from given name
	Cookie(name string) string
	// Cookies returns all map of cookies
	Cookies() map[string]string
	// Body returns body of response object
	Body() []byte
	// Parse converts response body to given object (JSON, XML)
	Parse(object any) error
	// FormFile returns content of file in bytes
	FormFile(key string) []byte

	// IsTLS returns flag is TLS
	IsTLS() bool
	// IsWebSocket returns flag is request websocket
	IsWebSocket() bool

	// Get returns context container variable
	Get(key string) any
	// Set creates new context container key-value pair
	Set(key string, value any)

	// Status sets HTTP status code to response
	Status(status int) Context

	// Empty writes empty response
	Empty() error
	// String writes string response
	String(message string) error
	// Bytes writes response body of given bytes array
	Bytes(body []byte) error
	// JSON writes response body of given object converted to JSON
	JSON(body any) error
	// XML writes response body of given object converted to XML
	XML(body any) error
	// Error writes response body of given error to JSON object
	Error(err error) error

	// Next method which calls next handler from handlers chain
	Next() error
}
