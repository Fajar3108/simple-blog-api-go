package router

import (
	"github.com/julienschmidt/httprouter"
	"simple-blog-api-golang/internal/database"
	"simple-blog-api-golang/pkg/di"
	"simple-blog-api-golang/pkg/jwt_service"
)

func SetupRouter() *httprouter.Router {
	router := httprouter.New()

	var db, err = database.InitDB()

	if err != nil {
		panic(err)
	}

	dependency, err := di.InitializeApp(db)

	if err != nil {
		panic(err)
	}

	router.POST("/api/auth/login", dependency.AuthHandler.Signin)
	router.POST("/api/auth/register", dependency.AuthHandler.Signup)

	router.GET("/api/posts", jwt_service.AuthMiddleware(dependency.PostHandler.Index))
	router.GET("/api/posts/:id", jwt_service.AuthMiddleware(dependency.PostHandler.Show))

	return router
}
