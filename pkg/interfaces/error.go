package interfaces

// Error is interface of Boost Error. It is custom error of Boost
type Error interface {
	HttpCode() int
	SetHttpCode(code int) Error
	Type() string
	SetType(errorType string) Error
	Error() string
	ContentType() string
	JSON() []byte
	Is(compare error) bool
}
