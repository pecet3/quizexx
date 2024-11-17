package database

import (
	"database/sql"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func ConnectDb() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file: ", err)
		return nil
	}
	user := os.Getenv("DB_USER")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")
	sslmode := os.Getenv("DB_SSLMODE")

	connStr := "user=" + user + " host=" + host + " port=" + port + " password=" + password + " dbname=" + dbname + " sslmode=" + sslmode
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("error during open sql: ", err)
		return nil
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("error during open sql: ", err)
		return nil
	}

	err = prepare(db)
	if err != nil {
		log.Println("config database error:", err)
		return nil
	}
	log.Println("> Created DB tables if did not exist")

	return db
}
