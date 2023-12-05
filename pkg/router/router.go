package router

import (
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func SetupRouter() *httprouter.Router {
	router := httprouter.New()

	router.GET("/", func(writer http.ResponseWriter, request *http.Request, _ httprouter.Params) {
		fmt.Fprintln(writer, "Hello World")
	})

	return router
}
