package router

import (
	"database/sql"
	"net/http"

	"github.com/margar-melkonyan/watch-later.git/internal/handler/controller"
	"github.com/margar-melkonyan/watch-later.git/internal/repository"
	service "github.com/margar-melkonyan/watch-later.git/internal/service/users"
)

func authRoutes(db *sql.DB) *http.ServeMux {
	userRepo := repository.NewUserRepository(db)
	_ = service.NewUserService(userRepo)
	auth := http.NewServeMux()
	authController := controller.NewAuthController(db)

	auth.HandleFunc("POST /sign-in", authController.SignIn)
	auth.HandleFunc("POST /sign-up", authController.SignUp)
	auth.HandleFunc("GET /current-user", authController.CurrentUser)
	auth.HandleFunc("POST /refresh-token", authController.RefreshToken)

	return auth
}
