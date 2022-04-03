package utils

import (
	"database/sql"
	"log"
	"os"
)

func DatabaseInit() (db *sql.DB) {
	dsn := os.Getenv("DSN")
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println("db connection failed")
	}
	return
}
