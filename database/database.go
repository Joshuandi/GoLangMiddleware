package database

import "database/sql"

const (
	Host     = "localhost"
	Port     = 5432
	User     = "postgres"
	Password = "admin"
	Dbname   = "db-go-sql"
)

var (
	Db  *sql.DB
	Err error
)
