package database

import (
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
)

func RunMigrations(migrationsPath string, databaseUrl string) {
	m, err := migrate.New("file://"+migrationsPath, "sqlite3://"+databaseUrl)

	if err != nil {
		log.Fatal(err)
	}

	if err = m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal(err)
	}
}
