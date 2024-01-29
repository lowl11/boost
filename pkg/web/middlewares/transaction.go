package middlewares

import (
	"github.com/jmoiron/sqlx"
	"github.com/lowl11/boost/data/interfaces"
	"github.com/lowl11/boost/pkg/system/container"
	"github.com/lowl11/boost/pkg/system/types"
	"github.com/lowl11/boost/storage"
)

func Transaction() types.MiddlewareFunc {
	return func(ctx interfaces.Context) error {
		nativeCtx := ctx.Context()

		nativeCtx = storage.MustBeginTransaction(nativeCtx, container.Type[sqlx.DB]("connection"))
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
