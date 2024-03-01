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
			return errors.
				New("Used Transaction() middleware but wasn't set connection").
				SetType("Storage_NoConnection").
				AddContext("from", "Transaction Middleware")
		}

		nativeCtx = storage.MustBeginTx(nativeCtx, connection)
		ctx.SetContext(nativeCtx)
		defer storage.MustRollbackTx(nativeCtx)

		err := ctx.Next()
		if err != nil {
			storage.MustRollbackTx(nativeCtx)
			return err
		}

		storage.MustCommitTx(nativeCtx)
		return nil
	}
}
