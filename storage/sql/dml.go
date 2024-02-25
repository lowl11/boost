package sql

import "context"

type SelectBuilder interface {
	Query

	Single(ctx context.Context, args ...any) Scanner
	List(ctx context.Context, args ...any) Scanner
	Scan(ctx context.Context, result any, args ...any) error
	ScanSingle(ctx context.Context, result any, args ...any) error

	Select(columns ...string) SelectBuilder
	Distinct() SelectBuilder
	From(tableName string) SelectBuilder
	As(aliasName string) SelectBuilder
	Join(tableName, aliasName, joinColumn, mainColumn string) SelectBuilder
	LeftJoin(tableName, aliasName, joinColumn, mainColumn string) SelectBuilder
	RightJoin(tableName, aliasName, joinColumn, mainColumn string) SelectBuilder
	Where(func(Where)) SelectBuilder
	OrderBy(columns ...string) SelectBuilder
	OrderByDescending(columns ...string) SelectBuilder
	Having(func(Aggregate)) SelectBuilder
	GroupBy(columns ...string) SelectBuilder
	GroupByAggregate(func(Aggregate)) SelectBuilder
	Offset(offset int) SelectBuilder
	Limit(limit int) SelectBuilder
	Page(pageSize, pageNumber int) SelectBuilder
}

type DeleteBuilder interface {
	Query

	Exec(ctx context.Context, args ...any) error

	From(tableName string) DeleteBuilder
	Where(func(Where)) DeleteBuilder
}

type UpdateBuilder interface {
	Query

	Exec(ctx context.Context) error

	GetParam() (string, bool)
	From(tableName string) UpdateBuilder
	Set(pairs ...Pair) UpdateBuilder
	Where(func(Where)) UpdateBuilder
	Entity(entity any) UpdateBuilder
}

type InsertBuilder interface {
	Query

	Exec(ctx context.Context) error

	GetParamStatus() (string, bool)
	To(tableName string) InsertBuilder
	OnConflict(query string) InsertBuilder
	Values(pairs ...Pair) InsertBuilder
	Entity(entity any) InsertBuilder
	EntityList(list []any) InsertBuilder
}

func Select(columns ...string) SelectBuilder {
	return newSelectBuilder(columns...)
}

func Delete(tableName ...string) DeleteBuilder {
	return newDeleteBuilder(tableName...)
}

func Update(tableName ...string) UpdateBuilder {
	return newUpdateBuilder(tableName...)
}

func Insert(pairs ...Pair) InsertBuilder {
	return newInsertBuilder(pairs...)
}
