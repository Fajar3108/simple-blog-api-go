package main

import (
	"fmt"
	"log"
	"net/http"
	"simple-blog-api-golang/internal/database"
	"simple-blog-api-golang/pkg/router"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := database.InitDB("root", "", "localhost", "3306", "simple_blog_api_go")
	defer database.CloseDB()

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(10)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	fmt.Println("Server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", router.SetupRouter())

	if err != nil {
		log.Fatal(err)
	}
}
