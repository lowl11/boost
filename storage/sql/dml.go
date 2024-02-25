package sql

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
