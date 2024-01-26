package context

import (
	"github.com/lowl11/boost/data/interfaces"
	"github.com/lowl11/boost/internal/boosties/errors"
	"github.com/lowl11/boost/internal/helpers/type_helper"
)

func (ctx *Context) returnOKObject(value any) error {
	if type_helper.IsPrimitive(value) {
		return ctx.String(type_helper.ToString(value, false))
	}

	return ctx.JSON(value)
}

func (ctx *Context) redirect(url string, customStatus ...int) error {
	if url == "" {
		return errors.New("Given Redirect URL is empty")
	}

	ctx.inner.Response.Header.SetCanonical(
		type_helper.StringToBytes("Location"),
		type_helper.StringToBytes(url),
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
