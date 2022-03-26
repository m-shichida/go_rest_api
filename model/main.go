package model

import (
	"database/sql"
	"log"
	"os"
)

var Db *sql.DB
var err error

func DatabaseInit() {
	dsn := os.Getenv("DSN")
	Db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Println("db connection failed")
	}
}
