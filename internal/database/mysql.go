package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

var db *gorm.DB

func InitDB() *gorm.DB {
	if db != nil {
		return db
	}

	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", user, password, host, port, dbName)

	var err error

	db, err = gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	}

	return db
}
