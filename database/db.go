package database

import (
	"database/sql"
	"log/slog"

	_ "github.com/lib/pq"
)

func DB() *sql.DB {
	connStr := "host=localhost port=5432 user=postgres password=postgres dbname=go sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		slog.Error("Error while connection to the db", err)
	}

	err = db.Ping()

	if err != nil {
		slog.Error("Cannot ping the db", err)
	}

	return db
}