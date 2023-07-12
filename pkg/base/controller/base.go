package controller

import (
	"github.com/lowl11/boost"
	"github.com/lowl11/boost/internal/controller/domain"
	"github.com/lowl11/boost/internal/helpers/type_helper"
	"github.com/lowl11/lazylog/log"
	"net/http"
)

type Base struct {
	WrappedOK bool
}

func (controller *Base) Ok(ctx boost.Context, body ...any) error {
	ctx.Status(http.StatusOK)
	if len(body) > 0 {
		if controller.WrappedOK {
			return ctx.JSON(domain.NewWrappedOK(body[0]))
		}

		return controller.returnOKObject(ctx, body[0])
	}

	return ctx.JSON(domain.NewJustOK())
}

func (controller *Base) Created(ctx boost.Context) error {
	return ctx.Status(http.StatusCreated).Empty()
}

func (controller *Base) CreatedBody(ctx boost.Context, body any) error {
	return ctx.Status(http.StatusCreated).JSON(body)
}

func (controller *Base) CreatedID(ctx boost.Context, id any) error {
	return ctx.Status(http.StatusCreated).JSON(domain.NewCreatedWithID(id))
}

func (controller *Base) NotFound(ctx boost.Context) error {
	return ctx.Status(http.StatusNotFound).Empty()
}

func (controller *Base) Error(ctx boost.Context, err error) error {
	if err == nil {
		return ctx.Status(http.StatusInternalServerError).Empty()
	}

	log.Error(err)
	return ctx.Error(err)
}

func (controller *Base) returnOKObject(ctx boost.Context, value any) error {
	if type_helper.IsPrimitive(value) {
		return ctx.String(type_helper.ToString(value))
	}

	return ctx.JSON(value)
}
