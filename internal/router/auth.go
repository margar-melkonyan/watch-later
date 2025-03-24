package router

import (
	"database/sql"
	"github.com/margar-melkonyan/watch-later.git/internal/handler/controller"
	"github.com/margar-melkonyan/watch-later.git/internal/repository"
	"github.com/margar-melkonyan/watch-later.git/internal/service/users"
	"net/http"
)

func authRoutes(db *sql.DB) *http.ServeMux {
	userRepo := repository.NewUserRepository(db)
	_ = service.NewUserService(userRepo)
	auth := http.NewServeMux()
	authController := controller.NewAuthController(db)

	auth.HandleFunc("POST /sign-in", authController.SignIn)
	auth.HandleFunc("POST /sign-up", authController.SignUp)
	auth.HandleFunc("GET /current-user", func(w http.ResponseWriter, r *http.Request) {})
	auth.HandleFunc("GET /sign-out", authController.SignOut)
	auth.HandleFunc("POST /refresh-token", func(w http.ResponseWriter, r *http.Request) {})

	return auth
}
