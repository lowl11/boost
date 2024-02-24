package sql

import (
	"context"
	"github.com/lowl11/boost/log"
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
			return nil
		}

		return rows.StructScan(result)
	}

	return repo.DB(s.ctx).SelectContext(s.ctx, result, s.query, s.args...)
}
