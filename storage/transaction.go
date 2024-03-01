package storage

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/lowl11/boost/log"
	"strings"
)

func ReadTx(ctx context.Context) *sqlx.Tx {
	if ctx == nil {
		return nil
	}

	txValue := ctx.Value("boost_transaction")
	if txValue == nil {
		return nil
	}

	return txValue.(*sqlx.Tx)
}

func BeginTx(ctx context.Context, connection *sqlx.DB) (context.Context, error) {
	tx, err := connection.BeginTxx(ctx, &sql.TxOptions{
		Isolation: sql.LevelDefault,
		ReadOnly:  false,
	})
	if err != nil {
		return nil, err
	}

	return context.WithValue(ctx, "boostef_transaction", tx), nil
}

func MustBeginTx(ctx context.Context, connection *sqlx.DB) context.Context {
	newCtx, err := BeginTx(ctx, connection)
	if err != nil {
		return ctx
	}

	return newCtx
}

func RollbackTx(ctx context.Context) error {
	tx := ReadTx(ctx)
	if tx == nil {
		return nil
	}

	if err := tx.Rollback(); err != nil {
		if strings.Contains(err.Error(), "transaction has already been committed") {
			return nil
		}

		return err
	}

	return nil
}

func MustRollbackTx(ctx context.Context) {
	if err := RollbackTx(ctx); err != nil {
		log.Error("Rollback transaction error:", err)
	}
}

func CommitTx(ctx context.Context) error {
	tx := ReadTx(ctx)
	if tx == nil {
		return nil
	}

	if err := tx.Commit(); err != nil {
		if strings.Contains(err.Error(), "transaction has already been committed") {
			return nil
		}

		return err
	}

	return nil
}

func MustCommitTx(ctx context.Context) {
	if err := CommitTx(ctx); err != nil {
		log.Error("Commit transaction error:", err)
	}
}
