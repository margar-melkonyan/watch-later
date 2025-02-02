package router

import (
	"database/sql"
	"net/http"
)

func platformsRoutes(db *sql.DB) *http.ServeMux {
	platforms := http.NewServeMux()

	platforms.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {})
	platforms.HandleFunc("GET /{id}", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(r.PathValue("id")))
	})
	platforms.HandleFunc("POST /", func(w http.ResponseWriter, r *http.Request) {})
	platforms.HandleFunc("POST /{id}/restore", func(w http.ResponseWriter, r *http.Request) {})
	platforms.HandleFunc("PUT /{id}", func(w http.ResponseWriter, r *http.Request) {})
	platforms.HandleFunc("DELETE /{id}", func(w http.ResponseWriter, r *http.Request) {})

	return platforms
}
