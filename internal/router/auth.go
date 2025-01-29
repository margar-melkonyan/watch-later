package router

import (
	"net/http"
)

func authRoutes() *http.ServeMux {
	auth := http.NewServeMux()

	auth.HandleFunc("POST /sign-in", func(w http.ResponseWriter, r *http.Request) {})
	auth.HandleFunc("POST /sign-up", func(w http.ResponseWriter, r *http.Request) {})
	auth.HandleFunc("GET /current-user", func(w http.ResponseWriter, r *http.Request) {})
	auth.HandleFunc("GET /sign-out", func(w http.ResponseWriter, r *http.Request) {})
	auth.HandleFunc("POST /refresh-token", func(w http.ResponseWriter, r *http.Request) {})

	return auth
}
