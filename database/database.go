package database

import (
	"database/sql"
	"os"

	_ "github.com/lib/pq"
)

func ConnectToDatabase() *sql.DB {
	uri := os.Getenv("DATABASE_URI")
	if uri == "" {
		panic("You must set your 'DATABASE_URI' environmental variable")
	}

	db, err := sql.Open("postgres", uri)
	if err != nil {
		panic(err.Error())
	}

	return db
}
