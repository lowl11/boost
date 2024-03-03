package interfaces

import (
	"context"
	"github.com/valyala/fasthttp"
	"io"
)

// Context is interface of context which comes to Boost Handler
type Context interface {
	// Request returns Fast HTTP request object
	Request() *fasthttp.Request
	// Response returns Fast HTTP response object
	Response() *fasthttp.Response

	// Method returns request method
	Method() string
	// Scheme returns scheme of request (http/https)
	Scheme() string

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
	// Authorization returns "Authorization" header value without
	Authorization() string
	// Body returns body of response object
	Body() []byte
	// Parse converts response body to given object (JSON, XML).
	// Using validation ("validate" tag by go-playground). Turn off validation in Boost Config
	Parse(object any) error
	// Validate check body fields by using "validate" tag
	Validate(object any) error
	// FormFile returns content of file in bytes
	FormFile(key string) []byte
	FormValue(key string) Param

	// IsTLS returns flag is TLS
	IsTLS() bool
	// IsWebSocket returns flag is request websocket
	IsWebSocket() bool

	// Get returns context container variable
	Get(key string) any
	// Set creates new context container key-value pair
	Set(key string, value any)

	// SetCookie sets new cookie key=value to response
	SetCookie(key, value string, opts ...func(cookie *fasthttp.Cookie)) Context
	// SetHeader sets new header key=value to response
	SetHeader(key, value string) Context

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
	// ThrowError writes response body of given error to JSON object
	ThrowError(err error) error
	// Redirect redirects to the given URL
	Redirect(url string) error
	// RedirectStatus redirects to the given URL and with given status
	RedirectStatus(url string, status ...int) error

	// Next method which calls next handler from handlers chain
	Next() error

	// Context returns context.Context.
	// If there was set timeout on Boost App config, context will be with timeout.
	// Also, there is possibility to add custom context at middleware level with SetContext method
	Context() context.Context
	// SetContext possibility to set custom context.Context and boost context will return it at Context() method
	SetContext(ctx context.Context)

	// SetPanicHandler sets custom panic handler function
	SetPanicHandler(panicHandler func(err error))
	// PanicHandler returns custom panic handler function
	PanicHandler() func(err error)

	// Ok returns response with code 200 and given body.
	// Note: if given body is primitive variable (int, string, bool, etc.) it will be returned with text/plain
	Ok(body ...any) error
	// Created returns response with code 201, with no body
	Created() error
	// CreatedBody returns response with code 201, with given body
	CreatedBody(body any) error
	// CreatedID returns response with code 201, with given ID to return.
	// Note: response object will be in JSON. If id is int, will be returns int, if string will be returned string
	// Example:
	//
	//	{
	//		"id": 123
	//	}
	CreatedID(id any) error
	// NotFound returns response with status 404, with no body
	NotFound() error
	// NotFoundError returns response with status 404, with given body
	NotFoundError(err error) error
	// NotFoundString returns response with status 404, with given message
	NotFoundString(message string) error
	// Error returns response with given error status, error object.
	// Note: if given err will not be defined as Boost Error, default status code is 500
	Error(err error) error

	Writer() io.Writer
}
