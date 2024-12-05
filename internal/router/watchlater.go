package router

import "net/http"

func watchLatersRoutes() *http.ServeMux {
	watchLaters := http.NewServeMux()

	watchLaters.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {})
	watchLaters.HandleFunc("GET /{id}", func(w http.ResponseWriter, r *http.Request) {})
	watchLaters.HandleFunc("POST /", func(w http.ResponseWriter, r *http.Request) {})
	watchLaters.HandleFunc("PUT /{id}", func(w http.ResponseWriter, r *http.Request) {})
	watchLaters.HandleFunc("DELETE /{id}", func(w http.ResponseWriter, r *http.Request) {})

	return watchLaters
}
