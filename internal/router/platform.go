package router

import "net/http"

func platofmrsRoutes() *http.ServeMux {
	platofmrs := http.NewServeMux()

	platofmrs.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {})
	platofmrs.HandleFunc("GET /{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(r.PathValue("id")))
	})
	platofmrs.HandleFunc("POST /", func(w http.ResponseWriter, r *http.Request) {})
	platofmrs.HandleFunc("POST /{id}/restore", func(w http.ResponseWriter, r *http.Request) {})
	platofmrs.HandleFunc("PUT /{id}", func(w http.ResponseWriter, r *http.Request) {})
	platofmrs.HandleFunc("DELETE /{id}", func(w http.ResponseWriter, r *http.Request) {})

	return platofmrs
}
