package data

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func NewSQLite() *sql.DB {
	db, err := sql.Open("sqlite3", "./store.db")
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	return db
}
