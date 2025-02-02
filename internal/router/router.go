package router

import (
	"database/sql"
	"net/http"
)

func NewRouter(db *sql.DB) *http.ServeMux {
	router := http.NewServeMux()
	api := http.NewServeMux()
	v1 := http.NewServeMux()

	v1.Handle("/auth/", http.StripPrefix("/auth", authRoutes(db)))
	v1.Handle("/notifications/", http.StripPrefix("/notifications", notificationsRoutes(db)))
	v1.Handle("/categories/", http.StripPrefix("/categories", categoriesRoutes(db)))
	v1.Handle("/platforms/", http.StripPrefix("/platforms", platformsRoutes(db)))
	v1.Handle("/watch-laters/", http.StripPrefix("/watch-laters", watchLatersRoutes(db)))
	// /api/v1/
	api.Handle("/v1/", http.StripPrefix("/v1", v1))
	// /api/
	router.Handle("/api/", http.StripPrefix("/api", api))

	return router
}
