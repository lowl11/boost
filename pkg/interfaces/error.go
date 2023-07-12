package interfaces

type Error interface {
	HttpCode() int
	SetHttpCode(code int) Error
	Type() string
	SetType(errorType string) Error
	Error() string
	ContentType() string
	JSON() []byte
}
