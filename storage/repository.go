package storage

import (
	"github.com/jmoiron/sqlx"
	"github.com/lowl11/boost/log"
	"github.com/lowl11/boost/pkg/system/di"
	"strings"
)

type Repository interface {
	CloseRows(rows *sqlx.Rows)
	Transaction(transactionActions func(tx *sqlx.Tx) error) error
}

type repository struct {
	Connection *sqlx.DB
}

func NewRepo() Repository {
	return repository{
		Connection: di.Get[sqlx.DB](),
	}
}

func (repo repository) CloseRows(rows *sqlx.Rows) {
	if err := rows.Close(); err != nil {
		log.Error("Closing rows error:", err)
	}
}

func (repo repository) Transaction(transactionActions func(tx *sqlx.Tx) error) error {
	transaction, err := repo.Connection.Beginx()
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
