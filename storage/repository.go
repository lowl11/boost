package storage

import (
	"context"
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/lowl11/boost/log"
	"github.com/lowl11/boost/pkg/system/di"
	"strings"
)

type Repository interface {
	CloseRows(rows *sqlx.Rows)
	Transaction(transactionActions func(tx *sqlx.Tx) error) error
	Connection() *sqlx.DB
	DB(ctx context.Context) DB
}

type DB interface {
	sqlx.ExecerContext
	sqlx.QueryerContext
	sqlx.PreparerContext
	NamedExecContext
	SelectContext
}

type NamedExecContext interface {
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
}

type SelectContext interface {
	SelectContext(ctx context.Context, dest interface{}, query string, args ...interface{}) error
}

type repository struct {
	connection *sqlx.DB
}

func NewRepo() Repository {
	return repository{
		connection: di.Get[sqlx.DB](),
	}
}

func (repo repository) Connection() *sqlx.DB {
	return repo.connection
}

func (repo repository) DB(ctx context.Context) DB {
	tx := getTransaction(ctx)
	if tx != nil {
		return tx
	}

	return repo.connection
}

func (repo repository) CloseRows(rows *sqlx.Rows) {
	if err := rows.Close(); err != nil {
		log.Error("Closing rows error:", err)
	}
}

func (repo repository) Transaction(transactionActions func(tx *sqlx.Tx) error) error {
	transaction, err := repo.connection.Beginx()
	if err != nil {
		return err
	}
	defer rollback(transaction)

	if err = transactionActions(transaction); err != nil {
		return err
	}

	if err = transaction.Commit(); err != nil {
		return err
	}

	return nil
}

func rollback(transaction *sqlx.Tx) {
	if err := transaction.Rollback(); err != nil {
		if !strings.Contains(
			err.Error(),
			"transaction has already been committed or rolled back",
		) {
			log.Error("Rollback transaction error:", err)
		}
	}
}
