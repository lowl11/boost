package sql

type SelectBuilder interface {
	Query

	Select(columns ...string) SelectBuilder
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

	From(tableName string) DeleteBuilder
	Where(func(Where)) DeleteBuilder
}

type UpdateBuilder interface {
	Query

	GetParam() (string, bool)
	From(tableName string) UpdateBuilder
	Set(pairs ...Pair) UpdateBuilder
	Where(func(Where)) UpdateBuilder
}

type InsertBuilder interface {
	Query

	GetParamStatus() (string, bool)
	To(tableName string) InsertBuilder
	OnConflict(query string) InsertBuilder
	Values(pairs ...Pair) InsertBuilder
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
