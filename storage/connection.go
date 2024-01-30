package storage

import (
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/lowl11/boost/pkg/system/container"
	"time"
)

func NewPool(connectionString string, options ...func(connection *sqlx.DB)) (*sqlx.DB, error) {
	pgxConfig, _ := pgx.ParseConfig(connectionString)

	connection, err := sqlx.Open("pgx", stdlib.RegisterConnConfig(pgxConfig))
	if err != nil {
		return nil, err
	}

	for _, option := range options {
		option(connection)
	}

	if err = connection.Ping(); err != nil {
		return nil, err
	}

	container.Set("connection", connection)
	return connection, nil
}

func WithMaxConnectionsOption(maxOpenConnections int) func(db *sqlx.DB) {
	return func(db *sqlx.DB) {
		db.SetMaxOpenConns(maxOpenConnections)
	}
}

func WithMaxIdleConnectionsOption(maxIdleConnections int) func(db *sqlx.DB) {
	return func(db *sqlx.DB) {
		db.SetMaxIdleConns(maxIdleConnections)
	}
}

func WithMaxLifetime(lifetime time.Duration) func(db *sqlx.DB) {
	return func(db *sqlx.DB) {
		db.SetConnMaxLifetime(lifetime)
	}
}

func WithMaxIdleLifetime(lifetime time.Duration) func(db *sqlx.DB) {
	return func(db *sqlx.DB) {
		db.SetConnMaxIdleTime(lifetime)
	}
}
