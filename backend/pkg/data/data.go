package data

import (
	"database/sql"
	"log"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

const PREFIX = "/v1"

func newDb() *sql.DB {
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatalf("Failed to connect database: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	prepareList := [2]string{
		UsersTable,
		SessionsTable,
	}

	for _, table := range prepareList {
		if err := prepare(db, table); err != nil {
			log.Fatalf("Failed to prepare DB: %v", err)
		}
	}

	log.Println("<DB> Preparing DB has been finished")

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
	log.Print("<DB> ", table[1:endIndex])

	return nil
}

type Data struct {
	Db   *sql.DB
	User User
}

func New() *Data {
	return &Data{
		Db:   newDb(),
		User: User{},
	}
}
