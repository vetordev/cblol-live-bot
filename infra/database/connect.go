package database

import (
	"database/sql"
	"log"
	"os"
)

func Connect(url string) *sql.DB {
	if _, err := os.Stat(url); os.IsNotExist(err) {
		os.Create(url)
	}
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
