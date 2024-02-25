package storage

import (
	"context"
	"github.com/lowl11/boost/storage/sql"
)

type Crud struct {
	table string
	alias string
}

func NewCrud(table, alias string) *Crud {
	return &Crud{
		table: table,
		alias: alias,
	}
}

func (crud Crud) Add(ctx context.Context, entity any) error {
	return sql.Insert().Entity(entity).Exec(ctx)
}

func (crud Crud) Update(ctx context.Context, entity any) error {
	return sql.Update().Entity(entity).Exec(ctx)
}
