package middlewares

import (
	"context"
	"github.com/lowl11/boost/data/domain"
	"github.com/lowl11/boost/data/interfaces"
	"github.com/lowl11/boost/errors"
	"github.com/lowl11/boost/pkg/io/exception"
	"net/http"
	"time"
)

func Timeout(timeout time.Duration) domain.MiddlewareFunc {
	return func(ctx interfaces.Context) error {
		if timeout == 0 {
			if err := ctx.Next(); err != nil {
				return err
			}

			return nil
		}

		ch := make(chan struct{}, 1)
		errorChannel := make(chan error, 1)
		timeoutCtx, cancel := context.WithTimeout(ctx.Context(), timeout)
		defer cancel()

		ctx.SetContext(timeoutCtx)

		go func() {
			defer func() {
				errorChannel <- exception.CatchPanic(recover())
			}()
			if err := ctx.Next(); err != nil {
				errorChannel <- err
			}
			ch <- struct{}{}
		}()

		select {
		case <-ch:
			return ctx.Next()
		case err := <-errorChannel:
			return ctx.Error(err)
		case <-timeoutCtx.Done():
			return errorTimeout()
		}
	}
}

func errorTimeout() error {
	const typeErrorTimeout = "Timeout"

	return errors.New("Request reached timeout!").
		SetType(typeErrorTimeout).
		SetHttpCode(http.StatusRequestTimeout)
}
