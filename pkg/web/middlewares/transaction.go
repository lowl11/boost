package middlewares

import (
	"github.com/jmoiron/sqlx"
	"github.com/lowl11/boost/data/interfaces"
	"github.com/lowl11/boost/errors"
	"github.com/lowl11/boost/pkg/system/di"
	"github.com/lowl11/boost/pkg/system/types"
	"github.com/lowl11/boost/storage"
)

func Transaction() types.MiddlewareFunc {
	return func(ctx interfaces.Context) error {
		nativeCtx := ctx.Context()

		connection := di.Get[sqlx.DB]()
		if connection == nil {
			return errors.New("Used Transaction() middleware but wasn't set connection")
		}

		nativeCtx = storage.MustBeginTransaction(nativeCtx, connection)
		ctx.SetContext(nativeCtx)
		defer storage.MustRollbackTransaction(nativeCtx)

		err := ctx.Next()
		if err != nil {
			storage.MustRollbackTransaction(nativeCtx)
			return err
		}

		storage.MustCommitTransaction(nativeCtx)
		return nil
	}
}
