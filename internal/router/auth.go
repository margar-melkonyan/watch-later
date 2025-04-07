package router

import (
	"database/sql"
	"net/http"

	"github.com/margar-melkonyan/watch-later.git/internal/handler/controller"
)

func authRoutes(db *sql.DB) *http.ServeMux {
	auth := http.NewServeMux()
	authController := controller.NewAuthController(db)

	auth.HandleFunc("POST /sign-in", authController.SignIn)
	auth.HandleFunc("POST /sign-up", authController.SignUp)
	auth.HandleFunc("GET /current-user", authController.CurrentUser)
	auth.HandleFunc("POST /refresh-token", authController.RefreshToken)

	return auth
}
