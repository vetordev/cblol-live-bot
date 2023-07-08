package database

import (
	"database/sql"
	"log"
)

func Connect(url string) *sql.DB {
	conn, err := sql.Open("sqlite3", url)

	if err != nil {
		log.Fatal(err)
	}

	err = conn.Ping()

	if err != nil {
		log.Fatal(err)
	}

	return conn
}
