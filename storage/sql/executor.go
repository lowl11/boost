package sql

import (
	"context"
	"github.com/lowl11/boost/log"
)

type Executor interface {
	Exec(args ...any) error
}

type executor struct {
	ctx    context.Context
	query  string
	entity any
}

func newExecutor(ctx context.Context, query string, entity any) *executor {
	return &executor{
		ctx:    ctx,
		query:  query,
		entity: entity,
	}
}

func newParamExecutor(ctx context.Context, query string) *executor {
	return &executor{
		ctx:   ctx,
		query: query,
	}
}

func Exec(ctx context.Context, query string, entity any, args ...any) error {
	return newExecutor(ctx, query, entity).Exec(args...)
}

func (e *executor) Exec(args ...any) error {
	repo := getRepo()

	if _logEnabled {
		log.Debug(e.query)
	}

	if e.entity == nil {
		_, err := repo.DB(e.ctx).ExecContext(e.ctx, e.query, args...)
		if err != nil {
			return err
		}
	} else {
		_, err := repo.DB(e.ctx).NamedExecContext(e.ctx, e.query, e.entity)
		if err != nil {
			return err
		}
	}

	return nil
}
