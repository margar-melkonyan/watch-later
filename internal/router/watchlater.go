package router

import (
	"database/sql"
	"github.com/margar-melkonyan/watch-later.git/internal/repository"
	"github.com/margar-melkonyan/watch-later.git/internal/service/watchlaters"
	"net/http"
)

func watchLatersRoutes(db *sql.DB) *http.ServeMux {
	watchLaters := http.NewServeMux()

	watchLaterRepo := repository.NewWatchLaterRepository(db)
	_ = service.NewWatchLaterService(watchLaterRepo)

	watchLaters.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) {})
	watchLaters.HandleFunc("GET /{id}", func(w http.ResponseWriter, r *http.Request) {})
	watchLaters.HandleFunc("POST /", func(w http.ResponseWriter, r *http.Request) {})
	watchLaters.HandleFunc("PUT /{id}", func(w http.ResponseWriter, r *http.Request) {})
	watchLaters.HandleFunc("DELETE /{id}", func(w http.ResponseWriter, r *http.Request) {})

	return watchLaters
}
