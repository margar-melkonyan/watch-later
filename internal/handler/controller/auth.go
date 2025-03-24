package controller

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/margar-melkonyan/watch-later.git/internal/helper"
	"github.com/margar-melkonyan/watch-later.git/internal/repository"
	service "github.com/margar-melkonyan/watch-later.git/internal/service/users"
)

type AuthController struct {
	authService *service.AuthService
}

func NewAuthController(db *sql.DB) *AuthController {
	userRepo := repository.NewUserRepository(db)
	return &AuthController{
		authService: service.NewAuthService(userRepo),
	}
}

func (a *AuthController) SignUp(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("content-type")
	if contentType != "" &&
		strings.ToLower(strings.TrimSpace(contentType)) != "application/json" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var userForm repository.User
	r.Body = http.MaxBytesReader(w, r.Body, 1024*1024*10)
	err := json.NewDecoder(r.Body).Decode(&userForm)
	if err != nil && err == io.EOF {
		json, err := json.Marshal(helper.Response{
			Data: "Заполните все необходимые поля",
		})
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		http.Error(w, string(json), http.StatusBadRequest)
		return
	}

	validate := validator.New()
	err = validate.Struct(userForm)

	if err != nil {
		http.Error(w, err.Error(), http.StatusUnprocessableEntity)
		return
	}

	ok := a.authService.SignUp(&userForm)

	if ok != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (a *AuthController) SignIn(w http.ResponseWriter, r *http.Request) {

}

func (a *AuthController) SignOut(w http.ResponseWriter, r *http.Request) {

}
