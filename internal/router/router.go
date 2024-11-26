package router

import (
	"net/http"
	"watch-later/internal/handler/controller"
)

func NewRouter() *http.ServeMux {
	router := http.NewServeMux()

	router.HandleFunc("GET /api/v1/hello-world", controller.HelloWorld)

	return router
}
