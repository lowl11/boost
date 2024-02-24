package sql

var (
	_sql string
)

func SetSQL(sql string) {
	_sql = sql
}

func getSQL() string {
	if _sql == "" {
		_sql = "Postgres"
	}
	return _sql
}
