package db

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var db *sql.DB

func Connection() {
	conn := "postgres://postgres:1234@localhost:5432/online_market_main?sslmode=disable"
	database, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}
	db = database
}

func GetDb() *sql.DB {
	return db
}
