package data

import (
	"database/sql"
	"log"
	"strings"

	"github.com/pecet3/quizex/pkg/logger"
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

func prepare(db *sql.DB, table string) error {
	s, err := db.Prepare(table)
	if err != nil {
		return err
	}
	_, err = s.Exec()
	if err != nil {
		return err
	}

	endIndex := strings.Index(table, "(")
	logger.InfoC(table[1:endIndex])

	return nil
}
