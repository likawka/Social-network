package sqlite

import (
	"database/sql"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type DBWrapper struct {
	DB *sql.DB
}

func NewDBWrapper() *DBWrapper {
	return &DBWrapper{}
}

func (dbw *DBWrapper) InitDB(DBPath, migrationsDir string) {
	var err error

	dbw.DB, err = sql.Open("sqlite3", DBPath)
	if err != nil {
		log.Fatal(err)
	}

	m, err := migrate.New(
		"file://"+migrationsDir,
		"sqlite3://"+DBPath,
	)
	if err != nil {
		log.Fatalf("migration initialization failed: %v", err)
	}

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatalf("migration failed: %v", err)
	} else if err == migrate.ErrNoChange {
		log.Println("No migrations to apply.")
	}
}

func (dbw *DBWrapper) Close() {
	if dbw.DB != nil {
		err := dbw.DB.Close()
		if err != nil {
			log.Fatalf("failed to close database: %v", err)
		}
	}
}
