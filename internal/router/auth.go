package router

import (
	"database/sql"
	"github.com/margar-melkonyan/watch-later.git/internal/repository"
	"github.com/margar-melkonyan/watch-later.git/internal/service/users"
	"net/http"
)

func authRoutes(db *sql.DB) *http.ServeMux {
	userRepo := repository.NewUserRepository(db)
	_ = service.NewUserService(userRepo)
	auth := http.NewServeMux()

	auth.HandleFunc("POST /sign-in", func(w http.ResponseWriter, r *http.Request) {})
	auth.HandleFunc("POST /sign-up", func(w http.ResponseWriter, r *http.Request) {})
	auth.HandleFunc("GET /current-user", func(w http.ResponseWriter, r *http.Request) {})
	auth.HandleFunc("GET /sign-out", func(w http.ResponseWriter, r *http.Request) {})
	auth.HandleFunc("POST /refresh-token", func(w http.ResponseWriter, r *http.Request) {})

	return auth
}
