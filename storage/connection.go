package storage

import (
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/lowl11/boost/errors"
	"github.com/lowl11/boost/log"
	"github.com/lowl11/boost/pkg/system/di"
	"time"
)

func Connect(connectionString string, options ...func(connection *sqlx.DB)) (*sqlx.DB, error) {
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

	return connection, nil
}

func MustConnect(connectionString string, options ...func(connection *sqlx.DB)) *sqlx.DB {
	connection, err := Connect(connectionString, options...)
	if err != nil {
		log.Fatal("Connect to Database error:", err)
	}

	return connection
}

func RegisterConnect(connectionString string, options ...func(connection *sqlx.DB)) {
	di.Register[sqlx.DB](func() *sqlx.DB {
		return MustConnect(connectionString, options...)
	})
}

func Ping() error {
	connection := di.Get[sqlx.DB]()
	if connection == nil {
		return errors.
			New("No database connection").
			SetType("Storage_NoConnection")
	}

	return connection.Ping()
}

func MustPing() {
	if err := Ping(); err != nil {
		panic(err)
	}
}

func Close(connection ...*sqlx.DB) {
	if len(connection) > 0 {
		if err := connection[0].Close(); err != nil {
			log.Error("Close Database connection error:", err)
		}

		return
	}

	containerConnection := di.Get[sqlx.DB]()
	if containerConnection == nil {
		return
	}

	Close(containerConnection)
}

func WithMaxConnections(maxOpenConnections int) func(db *sqlx.DB) {
	return func(db *sqlx.DB) {
		db.SetMaxOpenConns(maxOpenConnections)
	}
}

func WithMaxIdleConnections(maxIdleConnections int) func(db *sqlx.DB) {
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
