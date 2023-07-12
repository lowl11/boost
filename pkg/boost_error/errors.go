package boost_error

type Error interface {
	HttpCode() int
	Type() string
	Error() string
	ContentType() string
	JSON() []byte
}
