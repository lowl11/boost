package controller

import (
	"github.com/lowl11/boost"
	"github.com/lowl11/boost/internal/controller/domain"
	"github.com/lowl11/boost/internal/helpers/type_helper"
	"github.com/lowl11/lazylog/log"
	"net/http"
)

// Base is base controller with easy use methods to work with REST
type Base struct {
	WrappedOK bool
}

// Ok returns response with code 200 and given body.
// Note: if given body is primitive variable (int, string, bool, etc.) it will be returned with text/plain
func (controller Base) Ok(ctx boost.Context, body ...any) error {
	ctx.Status(http.StatusOK)
	if len(body) > 0 {
		if controller.WrappedOK {
			return ctx.JSON(domain.NewWrappedOK(body[0]))
		}

		return controller.returnOKObject(ctx, body[0])
	}

	return ctx.JSON(domain.NewJustOK())
}

// Created returns response with code 201, with no body
func (controller Base) Created(ctx boost.Context) error {
	return ctx.Status(http.StatusCreated).Empty()
}

// CreatedBody returns response with code 201, with given body
func (controller Base) CreatedBody(ctx boost.Context, body any) error {
	return ctx.Status(http.StatusCreated).JSON(body)
}

// CreatedID returns response with code 201, with given ID to return.
// Note: response object will be in JSON. If id is int, will be returns int, if string will be returned string
// Example:
//
//	{
//		"id": 123
//	}
func (controller Base) CreatedID(ctx boost.Context, id any) error {
	return ctx.Status(http.StatusCreated).JSON(domain.NewCreatedWithID(id))
}

// NotFound returns response with status 404, with no body
func (controller Base) NotFound(ctx boost.Context) error {
	return ctx.Status(http.StatusNotFound).Empty()
}

// NotFoundError returns response with status 404, with given body
func (controller Base) NotFoundError(ctx boost.Context, err error) error {
	return ctx.Status(http.StatusNotFound).Error(err)
}

// NotFoundString returns response with status 404, with given message
func (controller Base) NotFoundString(ctx boost.Context, message string) error {
	return ctx.Status(http.StatusNotFound).JSON(domain.NewNotFoundMessage(message))
}

// Error returns response with given error status, error object.
// Note: if given err will not be defined as Boost Error, default status code is 500
func (controller Base) Error(ctx boost.Context, err error) error {
	if err == nil {
		return ctx.Status(http.StatusInternalServerError).Empty()
	}

	log.Error(err)
	return ctx.Error(err)
}

func (controller Base) Redirect(ctx boost.Context, url string) error {
	return ctx.
		Status(http.StatusTemporaryRedirect).
		Redirect(url)
}

func (controller Base) returnOKObject(ctx boost.Context, value any) error {
	if type_helper.IsPrimitive(value) {
		return ctx.String(type_helper.ToString(value, false))
	}

	return ctx.JSON(value)
}
