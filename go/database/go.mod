module database_

go 1.15

replace (
	base => ../base
	worker => ../worker
)

require (
	base v0.0.0-00010101000000-000000000000
	github.com/go-sql-driver/mysql v1.5.0
	github.com/jmoiron/sqlx v1.3.4
	worker v0.0.0-00010101000000-000000000000
)
