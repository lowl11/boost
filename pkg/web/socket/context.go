package socket

import (
	"encoding/json"
	"github.com/lowl11/boost/pkg/system/types"
)

type Context struct {
	conn        *Conn
	messageType int
	body        []byte
}

func newContext(conn *Conn, messageType int, body []byte) *Context {
	return &Context{
		conn:        conn,
		messageType: messageType,
		body:        body,
	}
}

func (ctx *Context) Body() []byte {
	return ctx.body
}

func (ctx *Context) String() string {
	return types.ToString(ctx.body)
}

func (ctx *Context) Parse(export any) error {
	return json.Unmarshal(ctx.body, &export)
}

func (ctx *Context) Send(body any) error {
	return ctx.conn.WriteMessage(ctx.messageType, types.ToBytes(body))
}
