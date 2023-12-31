package middlewares

import (
	"context"
	"github.com/lowl11/boost/data/interfaces"
	"github.com/lowl11/boost/internal/boosties/errors"
	"github.com/lowl11/boost/internal/boosties/panicer"
	"github.com/lowl11/boost/pkg/system/types"
	"net/http"
	"time"
)

func Timeout(timeout time.Duration) types.MiddlewareFunc {
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
				errorChannel <- panicer.Handle(recover())
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
