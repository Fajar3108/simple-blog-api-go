package router

import (
	"github.com/julienschmidt/httprouter"
	"simple-blog-api-golang/internal/database"
	"simple-blog-api-golang/pkg/di"
)

func SetupRouter() *httprouter.Router {
	router := httprouter.New()

	db, err := database.InitDB()

	if err != nil {
		panic(err)
	}

	dependency, err := di.InitializeApp(db)

	if err != nil {
		panic(err)
	}

	router.GET("/api/posts", dependency.PostHandler.Index)
	router.GET("/api/posts/:id", dependency.PostHandler.Show)

	return router
}
