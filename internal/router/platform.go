package router

import (
	"database/sql"
	"net/http"

	"github.com/margar-melkonyan/watch-later.git/internal/handler/controller"
)

func platformsRoutes(db *sql.DB) *http.ServeMux {
	platforms := http.NewServeMux()

	platformController := controller.NewPlatformController(db)

	platforms.HandleFunc("GET /", platformController.GetPlatforms)
	platforms.HandleFunc("GET /{id}", platformController.GetPlatform)
	platforms.HandleFunc("POST /", platformController.StorePlatform)
	platforms.HandleFunc("POST /{id}/restore", platformController.RestorePlatform)
	platforms.HandleFunc("PUT /{id}", platformController.UpdatePlatform)
	platforms.HandleFunc("DELETE /{id}", platformController.DeletePlatform)

	return platforms
}
