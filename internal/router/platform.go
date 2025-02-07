package router

import (
	"database/sql"
	"github.com/margar-melkonyan/watch-later.git/internal/repository"
	service "github.com/margar-melkonyan/watch-later.git/internal/service/platforms"
	"net/http"
)

func platformsRoutes(db *sql.DB) *http.ServeMux {
	platforms := http.NewServeMux()
	platformRepo := repository.NewPlatformRepository(db)
	_ = service.NewPlatformService(platformRepo)

	platforms.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {})
	platforms.HandleFunc("GET /{id}", func(w http.ResponseWriter, r *http.Request) {})
	platforms.HandleFunc("POST /", func(w http.ResponseWriter, r *http.Request) {})
	platforms.HandleFunc("POST /{id}/restore", func(w http.ResponseWriter, r *http.Request) {})
	platforms.HandleFunc("PUT /{id}", func(w http.ResponseWriter, r *http.Request) {})
	platforms.HandleFunc("DELETE /{id}", func(w http.ResponseWriter, r *http.Request) {})

	return platforms
}
