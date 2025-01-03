package router

import "net/http"

func categoriesRoutes() *http.ServeMux {
	categories := http.NewServeMux()

	categories.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {})
	categories.HandleFunc("GET /{id}", func(w http.ResponseWriter, r *http.Request) {})
	categories.HandleFunc("POST /", func(w http.ResponseWriter, r *http.Request) {})
	categories.HandleFunc("POST /{id}/restore", func(w http.ResponseWriter, r *http.Request) {})
	categories.HandleFunc("PUT /{id}", func(w http.ResponseWriter, r *http.Request) {})
	categories.HandleFunc("DELETE /{id}", func(w http.ResponseWriter, r *http.Request) {})

	return categories
}
