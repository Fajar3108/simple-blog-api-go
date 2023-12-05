package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"simple-blog-api-golang/internal/database"
	"simple-blog-api-golang/internal/post"
	"simple-blog-api-golang/pkg/router"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal(err)
	}

	db := database.InitDB()

	err = db.AutoMigrate(&post.Post{})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Server running on http://localhost:8080")
	err = http.ListenAndServe(":8080", router.SetupRouter())

	if err != nil {
		log.Fatal(err)
	}
}
