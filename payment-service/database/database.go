package database

import (
	"database/sql"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// Database connection pool.
var PoolConnDB *sql.DB

// SetupDatabase
func SetupDatabase() {
	var err error
	PoolConnDB, err = sql.Open("sqlite3", "./paymentdb.db")
	if err != nil {
		log.Fatal("Unable to use data source name", err)
	}
	PoolConnDB.SetConnMaxLifetime(60 * time.Second)
	PoolConnDB.SetMaxIdleConns(3)
	PoolConnDB.SetMaxOpenConns(3)

	log.Println("Pool DB created")
}
