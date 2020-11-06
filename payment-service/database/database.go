package database

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var DbConn *sql.DB

// SetupDatabase
func SetupDatabase() {
	var err error
	DbConn, err = sql.Open("sqlite3", "./paymentdb.db")
	if err != nil {
		log.Println("TEST TESTSETE")
		log.Fatal(err)
	}
	DbConn.SetMaxOpenConns(3)
	DbConn.SetMaxIdleConns(3)
	DbConn.SetConnMaxLifetime(60 * time.Second)
}
