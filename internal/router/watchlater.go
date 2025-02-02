package router

import (
	"database/sql"
	"net/http"
)

func watchLatersRoutes(db *sql.DB) *http.ServeMux {
	watchLaters := http.NewServeMux()

	watchLaters.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {})
	watchLaters.HandleFunc("GET /{id}", func(w http.ResponseWriter, r *http.Request) {})
	watchLaters.HandleFunc("POST /", func(w http.ResponseWriter, r *http.Request) {})
	watchLaters.HandleFunc("PUT /{id}", func(w http.ResponseWriter, r *http.Request) {})
	watchLaters.HandleFunc("DELETE /{id}", func(w http.ResponseWriter, r *http.Request) {})

	return watchLaters
}
