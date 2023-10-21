package db

import (
	"database/sql"
	"log"
)

var db *sql.DB

func Connection() {
	conn := "postgres://postgres:1234@localhost:5432/consultation?sslmode=disable"
	database, err := sql.Open("postgres", conn)
	if err != nil {
		log.Fatal(err)
	}
	db = database
}

func GetDb() *sql.DB {
	return db
}
