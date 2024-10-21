package db

import (
	"database/sql"
	"log"
	"os"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDb() *sql.DB {
	if db != nil {
		return db
	}

	db, err := sql.Open("sqlite3", "data/db.sqlite3")
	if err != nil {
		log.Fatal(err)
	}

	txn, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}

	migrations, err := filepath.Glob("data/migrations/*.sql")
	if err != nil {
		log.Fatal(err)
	}
	for _, migrationPath := range migrations {
		log.Printf("Running migration %s", migrationPath)
		migration, err := os.ReadFile(migrationPath)
		if err != nil {
			txn.Rollback()
			log.Fatal(err)
		}

		_, err = txn.Exec(string(migration))
		if err != nil {
			txn.Rollback()
			log.Fatal(err)
		}
	}

	err = txn.Commit()
	if err != nil {
		log.Fatal(err)
	}

	log.Println("DB is ready.")

	return db
}
