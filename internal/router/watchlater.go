package router

import (
	"database/sql"
	"net/http"

	"github.com/margar-melkonyan/watch-later.git/internal/handler/controller"
)

func watchLatersRoutes(db *sql.DB) *http.ServeMux {
	watchLaters := http.NewServeMux()
	watchLaterController := controller.NewWatchLaterController(db)
	watchLaters.HandleFunc("GET /", watchLaterController.GetWatchLaters)
	watchLaters.HandleFunc("GET /category/{cateogry_id}", watchLaterController.GetWatchLatersByCategory)
	watchLaters.HandleFunc("GET /platform/{platform_id}", watchLaterController.GetWatchLatersByPlatform)
	watchLaters.HandleFunc("GET /{id}", watchLaterController.GetWatchLater)
	watchLaters.HandleFunc("POST /", watchLaterController.StoreWatchLater)
	watchLaters.HandleFunc("PUT /{id}", watchLaterController.UpdateWatchLater)
	watchLaters.HandleFunc("DELETE /{id}", watchLaterController.DeleteWatchLater)

	return watchLaters
}
