package sql

import (
	"context"
	"github.com/lowl11/boost/errors"
	"github.com/lowl11/boost/log"
	"github.com/lowl11/flex"
	"net/http"
	"reflect"
)

var (
	RecordNotFound = errors.
		New("record not found").
		SetHttpCode(http.StatusNotFound).
		SetType("Storage_RecordNotFound")
)

var (
	_logEnabled = false
)

func EnableLog() {
	_logEnabled = true
}

type Scanner interface {
	Scan(any) error
}

func Scan(ctx context.Context, query string, result any, args ...any) error {
	return newScanner(query, ctx, args...).Scan(result)
}

func ScanSingle(ctx context.Context, query string, result any, args ...any) error {
	return newScanner(query, ctx, args...).Single().Scan(result)
}

type scanner struct {
	ctx      context.Context
	args     []any
	query    string
	isSingle bool
}

func newScanner(query string, ctx context.Context, args ...any) *scanner {
	return &scanner{
		query: query,
		ctx:   ctx,
		args:  args,
	}
}

func (s *scanner) Single() *scanner {
	s.isSingle = true
	return s
}

func (s *scanner) Scan(result any) error {
	if _logEnabled {
		log.Debug(s.query)
	}

	repo := getRepo()

	if s.isSingle {
		rows, err := repo.DB(s.ctx).QueryxContext(s.ctx, s.query, s.args...)
		if err != nil {
			return err
		}
		defer repo.CloseRows(rows)

		if !rows.Next() {
			return RecordNotFound
		}

		if flex.Type(flex.Type(reflect.TypeOf(result)).Unwrap()).IsPrimitive() {
			return rows.Scan(result)
		}

		return rows.StructScan(result)
	}

	return repo.DB(s.ctx).SelectContext(s.ctx, result, s.query, s.args...)
}
