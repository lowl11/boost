package event_context

import "encoding/json"

func (ctx *Context) Body() []byte {
	return ctx.body
}

func (ctx *Context) Parse(object any) error {
	return json.Unmarshal(ctx.body, &object)
}
