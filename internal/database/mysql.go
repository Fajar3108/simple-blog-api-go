package database

import (
	"database/sql"
	"fmt"
	"log"
)

var db *sql.DB

func InitDB(user, password, host, port, dbName string) *sql.DB {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, dbName)

	var err error

	db, err = sql.Open("mysql", dataSourceName)

	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func CloseDB() {
	if db != nil {
		db.Close()
	}
}
