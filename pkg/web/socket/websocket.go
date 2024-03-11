package socket

import (
	"errors"
	"github.com/lowl11/boost/data/interfaces"
	"github.com/lowl11/boost/log"
	"github.com/lowl11/boost/pkg/io/exception"
	"github.com/lowl11/boost/pkg/system/types"
	"io"
	"sync"
	"time"

	"github.com/fasthttp/websocket"
	"github.com/valyala/fasthttp"
)

// Config ...
type Config struct {
	// Filter defines a function to skip middleware.
	// Optional. Default: nil
	Filter func(interfaces.Context) bool

	// HandshakeTimeout specifies the duration for the handshake to complete.
	HandshakeTimeout time.Duration

	// Subprotocols specifies the client's requested subprotocols.
	Subprotocols []string

	// Allowed Origin's based on the Origin header, this validate the request origin to
	// prevent cross-site request forgery. Everything is allowed if left empty.
	Origins []string

	// ReadBufferSize and WriteBufferSize specify I/O buffer sizes in bytes. If a buffer
	// size is zero, then a useful default size is used. The I/O buffer sizes
	// do not limit the size of the messages that can be sent or received.
	ReadBufferSize, WriteBufferSize int

	// WriteBufferPool is a pool of buffers for write operations. If the value
	// is not set, then write buffers are allocated to the connection for the
	// lifetime of the connection.
	//
	// A pool is most useful when the application has a modest volume of writes
	// across a large number of connections.
	//
	// Applications should use a single pool for each unique value of
	// WriteBufferSize.
	WriteBufferPool websocket.BufferPool

	// EnableCompression specifies if the client should attempt to negotiate
	// per message compression (RFC 7692). Setting this value to true does not
	// guarantee that compression will be supported. Currently only "no context
	// takeover" modes are supported.
	EnableCompression bool

	// RecoverHandler is a panic handler function that recovers from panics
	// Default recover function is used when nil and writes error message in a response field `error`
	// It prints stack trace to the stderr by default
	// Optional. Default: defaultRecover
	RecoverHandler func(*Conn)
}

func defaultRecover(c *Conn) {
	if err := exception.CatchPanic(recover()); err != nil {
		log.Error("Catch websocket panic:", err)
		_ = c.WriteJSON(err)
	}
}

// New returns a new `handler func(*Conn)` that upgrades a client to the
// websocket protocol, you can pass an optional config.
func New(handler func(*Conn), config ...Config) types.HandlerFunc {
	// Init config
	var cfg Config
	if len(config) > 0 {
		cfg = config[0]
	}
	if len(cfg.Origins) == 0 {
		cfg.Origins = []string{"*"}
	}
	if cfg.ReadBufferSize == 0 {
		cfg.ReadBufferSize = 1024
	}
	if cfg.WriteBufferSize == 0 {
		cfg.WriteBufferSize = 1024
	}
	if cfg.RecoverHandler == nil {
		cfg.RecoverHandler = defaultRecover
	}

	var upgrader = websocket.FastHTTPUpgrader{
		HandshakeTimeout:  cfg.HandshakeTimeout,
		Subprotocols:      cfg.Subprotocols,
		ReadBufferSize:    cfg.ReadBufferSize,
		WriteBufferSize:   cfg.WriteBufferSize,
		EnableCompression: cfg.EnableCompression,
		WriteBufferPool:   cfg.WriteBufferPool,
		CheckOrigin: func(fctx *fasthttp.RequestCtx) bool {
			if cfg.Origins[0] == "*" {
				return true
			}
			origin := types.ToString(fctx.Request.Header.Peek("Origin"))
			for i := range cfg.Origins {
				if cfg.Origins[i] == origin {
					return true
				}
			}
			return false
		},
	}

	return func(ctx interfaces.Context) error {
		if cfg.Filter != nil && !cfg.Filter(ctx) {
			return ctx.Next()
		}

		conn := acquireConn()
		// locals
		ctx.FastHttpContext().VisitUserValues(func(key []byte, value interface{}) {
			conn.locals[string(key)] = value
		})

		// params
		for _, param := range ctx.Params() {
			conn.params[param.String()] = ctx.Param(param.String()).String()

		}

		// queries
		ctx.FastHttpContext().QueryArgs().VisitAll(func(key, value []byte) {
			conn.queries[string(key)] = string(value)
		})

		// cookies
		ctx.FastHttpContext().Request.Header.VisitAllCookie(func(key, value []byte) {
			conn.cookies[string(key)] = string(value)
		})

		// headers
		ctx.FastHttpContext().Request.Header.VisitAll(func(key, value []byte) {
			conn.headers[string(key)] = string(value)
		})

		// ip address
		conn.ip = ctx.FastHttpContext().RemoteIP().String()

		if err := upgrader.Upgrade(ctx.FastHttpContext(), func(fconn *websocket.Conn) {
			conn.Conn = fconn
			defer releaseConn(conn)
			defer cfg.RecoverHandler(conn)
			handler(conn)
		}); err != nil { // Upgrading required
			return errors.New("upgrade required")
		}

		return nil
	}
}

// Conn https://godoc.org/github.com/gorilla/websocket#pkg-index
type Conn struct {
	*websocket.Conn
	locals  map[string]interface{}
	params  map[string]string
	cookies map[string]string
	headers map[string]string
	queries map[string]string
	ip      string
}

// Conn pool
var poolConn = sync.Pool{
	New: func() interface{} {
		return new(Conn)
	},
}

// Acquire Conn from pool
func acquireConn() *Conn {
	conn := poolConn.Get().(*Conn)
	conn.locals = make(map[string]interface{})
	conn.params = make(map[string]string)
	conn.queries = make(map[string]string)
	conn.cookies = make(map[string]string)
	conn.headers = make(map[string]string)
	return conn
}

// Return Conn to pool
func releaseConn(conn *Conn) {
	conn.Conn = nil
	poolConn.Put(conn)
}

// Locals makes it possible to pass interface{} values under string keys scoped to the request
// and therefore available to all following routes that match the request.
func (conn *Conn) Locals(key string, value ...interface{}) interface{} {
	if len(value) == 0 {
		return conn.locals[key]
	}
	conn.locals[key] = value[0]
	return value[0]
}

// Params is used to get the route parameters.
// Defaults to empty string "" if the param doesn't exist.
// If a default value is given, it will return that value if the param doesn't exist.
func (conn *Conn) Params(key string, defaultValue ...string) string {
	v, ok := conn.params[key]
	if !ok && len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return v
}

// Query returns the query string parameter in the url.
// Defaults to empty string "" if the query doesn't exist.
// If a default value is given, it will return that value if the query doesn't exist.
func (conn *Conn) Query(key string, defaultValue ...string) string {
	v, ok := conn.queries[key]
	if !ok && len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return v
}

// Cookies is used for getting a cookie value by key
// Defaults to empty string "" if the cookie doesn't exist.
// If a default value is given, it will return that value if the cookie doesn't exist.
func (conn *Conn) Cookies(key string, defaultValue ...string) string {
	v, ok := conn.cookies[key]
	if !ok && len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return v
}

// Headers is used for getting a header value by key
// Defaults to empty string "" if the header doesn't exist.
// If a default value is given, it will return that value if the header doesn't exist.
func (conn *Conn) Headers(key string, defaultValue ...string) string {
	v, ok := conn.headers[key]
	if !ok && len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return v
}

// IP returns the client's network address
func (conn *Conn) IP() string {
	return conn.ip
}

// Constants are taken from https://github.com/fasthttp/websocket/blob/master/conn.go#L43

// Close codes defined in RFC 6455, section 11.7.
const (
	CloseNormalClosure           = 1000
	CloseGoingAway               = 1001
	CloseProtocolError           = 1002
	CloseUnsupportedData         = 1003
	CloseNoStatusReceived        = 1005
	CloseAbnormalClosure         = 1006
	CloseInvalidFramePayloadData = 1007
	ClosePolicyViolation         = 1008
	CloseMessageTooBig           = 1009
	CloseMandatoryExtension      = 1010
	CloseInternalServerErr       = 1011
	CloseServiceRestart          = 1012
	CloseTryAgainLater           = 1013
	CloseTLSHandshake            = 1015
)

// The message types are defined in RFC 6455, section 11.8.
const (
	// TextMessage denotes a text data message. The text message payload is
	// interpreted as UTF-8 encoded text data.
	TextMessage = 1

	// BinaryMessage denotes a binary data message.
	BinaryMessage = 2

	// CloseMessage denotes a close control message. The optional message
	// payload contains a numeric code and text. Use the FormatCloseMessage
	// function to format a close message payload.
	CloseMessage = 8

	// PingMessage denotes a ping control message. The optional message payload
	// is UTF-8 encoded text.
	PingMessage = 9

	// PongMessage denotes a pong control message. The optional message payload
	// is UTF-8 encoded text.
	PongMessage = 10
)

var (
	// ErrBadHandshake is returned when the server response to opening handshake is
	// invalid.
	ErrBadHandshake = errors.New("websocket: bad handshake")
	// ErrCloseSent is returned when the application writes a message to the
	// connection after sending a close message.
	ErrCloseSent = errors.New("websocket: close sent")
	// ErrReadLimit is returned when reading a message that is larger than the
	// read limit set for the connection.
	ErrReadLimit = errors.New("websocket: read limit exceeded")
)

// FormatCloseMessage formats closeCode and text as a WebSocket close message.
// An empty message is returned for code CloseNoStatusReceived.
func FormatCloseMessage(closeCode int, text string) []byte {
	return websocket.FormatCloseMessage(closeCode, text)
}

// IsCloseError returns boolean indicating whether the error is a *CloseError
// with one of the specified codes.
func IsCloseError(err error, codes ...int) bool {
	return websocket.IsCloseError(err, codes...)
}

// IsUnexpectedCloseError returns boolean indicating whether the error is a
// *CloseError with a code not in the list of expected codes.
func IsUnexpectedCloseError(err error, expectedCodes ...int) bool {
	return websocket.IsUnexpectedCloseError(err, expectedCodes...)
}

// IsWebSocketUpgrade returns true if the client requested upgrade to the
// WebSocket protocol.
func IsWebSocketUpgrade(c interfaces.Context) bool {
	return websocket.FastHTTPIsWebSocketUpgrade(c.FastHttpContext())
}

// JoinMessages concatenates received messages to create a single io.Reader.
// The string term is appended to each message. The returned reader does not
// support concurrent calls to the Read method.
func JoinMessages(c *websocket.Conn, term string) io.Reader {
	return websocket.JoinMessages(c, term)
}
