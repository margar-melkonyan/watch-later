package router

import (
	"net/http"
)

func NewRouter() *http.ServeMux {
	router := http.NewServeMux()
	api := http.NewServeMux()
	v1 := http.NewServeMux()

	v1.Handle("/auth", http.StripPrefix("/auth", authRoutes()))
	v1.Handle("/notifications/", http.StripPrefix("/notifications", notificationsRoutes()))
	v1.Handle("/categories/", http.StripPrefix("/categories", categoriesRoutes()))
	v1.Handle("/platforms/", http.StripPrefix("/platforms", platofmrsRoutes()))
	v1.Handle("/watch-laters/", http.StripPrefix("/watch-laters", watchLatersRoutes()))
	// /api/v1/
	api.Handle("/v1/", http.StripPrefix("/v1", v1))
	// /api/
	router.Handle("/api/", http.StripPrefix("/api", api))

	return router
}
