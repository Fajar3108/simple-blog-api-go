package main

import (
	"context"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"net/http"
	"os"
	"os/signal"
	_ "simple-blog-api-golang/configs"
	"simple-blog-api-golang/internal/database"
	"simple-blog-api-golang/internal/post"
	"simple-blog-api-golang/internal/user"
	"simple-blog-api-golang/pkg/router"
	"syscall"
)

func main() {
	db, err := database.InitDB()

	if err != nil {
		log.Fatalf("error initializing database: %s", err)
	}

	if err := db.AutoMigrate(&post.Post{}, &user.User{}); err != nil {
		log.Fatal("Error auto migrating database:", err)
	}

	startServer()
}

func startServer() {
	fmt.Println("Server running on http://localhost:8080")

	server := http.Server{
		Addr:    ":8080",
		Handler: router.SetupRouter(),
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("Error starting server:", err)
		}
	}()

	waitForShutdown(&server)
}

func waitForShutdown(server *http.Server) {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	fmt.Println("Shutting down the server...")

	if err := server.Shutdown(context.Background()); err != nil {
		log.Println("Error shutting down the server:", err)
	}
}
