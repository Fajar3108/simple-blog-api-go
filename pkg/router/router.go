package router

import (
	"github.com/julienschmidt/httprouter"
	"simple-blog-api-golang/pkg/di"
)

func SetupRouter() *httprouter.Router {
	router := httprouter.New()

	postHandler := di.InitializeApp()

	router.GET("/api/posts", postHandler.Index)

	return router
}
