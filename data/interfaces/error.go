package interfaces

import "google.golang.org/grpc/codes"

// Error is interface of Boost Error. It is custom error of Boost
type Error interface {
	HttpCode() int
	GrpcCode() codes.Code
	SetHttpCode(code int) Error
	Type() string
	Message() string
	SetType(errorType string) Error
	Context() map[string]any
	SetContext(context map[string]any) Error
	AddContext(key string, value any) Error
	InnerError() error
	SetError(err error) Error
	Error() string
	String() string
	ContentType() string
	JSON() []byte
	Is(compare error) bool
}
