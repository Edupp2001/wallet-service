package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func NewPostgres() (*sql.DB, error) {
	dsn := "postgres://wallet:wallet@localhost:5432/wallet?sslmode=disable"
	return sql.Open("postgres", dsn)
}
