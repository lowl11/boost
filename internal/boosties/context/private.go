package context

import (
	"github.com/lowl11/boost/data/interfaces"
	"github.com/lowl11/boost/errors"
	"github.com/lowl11/boost/pkg/system/types"
)

func (ctx *Context) returnOKObject(value any) error {
	if types.IsPrimitive(value) {
		return ctx.String(types.ToString(value))
	}

	return ctx.JSON(value)
}

func (ctx *Context) redirect(url string, customStatus ...int) error {
	if url == "" {
		return errors.New("Given Redirect URL is empty")
	}

	ctx.inner.Response.Header.SetCanonical(
		types.StringToBytes("Location"),
		types.StringToBytes(url),
	)

	if len(customStatus) > 0 {
		ctx.Status(customStatus[0])
	}

	ctx.inner.Redirect(url, ctx.status)

	return nil
}

func (ctx *Context) returnError(err error) error {
	if err == nil {
		return nil
	}

	boostError, ok := err.(interfaces.Error)
	if !ok {
		boostError = ErrorUnknownType(err)
	}

	ctx.writer.Write(
		boostError.ContentType(),
		boostError.HttpCode(),
		boostError.JSON(),
	)

	return err
}
