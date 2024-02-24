package sql

func CreateTable(tableName ...string) CreateTableBuilder {
	return newCreateTable(tableName...)
}

func AlterTable(tableName ...string) AlterTableBuilder {
	return newAlterTableBuilder(tableName...)
}

func DropTable(tableName ...string) DropTableBuilder {
	return newDropTableBuilder(tableName...)
}

func TruncateTable(tableName ...string) TruncateTableBuilder {
	return newTruncateTableBuilder(tableName...)
}

func CreateIndex(name ...string) CreateIndexBuilder {
	return newCreateIndexBuilder(name...)
}

func DropIndex(name ...string) DropIndexBuilder {
	return newDropIndexBuilder(name...)
}
