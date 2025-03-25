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
		errs := err.(validator.ValidationErrors)

		var customMessages = map[string]string{
			"required": "The {field} field is required.",
			"email":    "The {field} must be a valid email address.",
			"min":      "The {field} must be at least {param} characters long.",
			"max":      "The {field} must be at most {param} characters long.",
			"gte":      "The {field} must be greater than or equal to {param}.",
			"lte":      "The {field} must be less than or equal to {param}.",
		}

		humanReadableError := make(map[string]interface{})

		for _, err := range errs {
			var res string
			res = strings.ReplaceAll(
				customMessages[err.Tag()],
				"{field}", strings.ToLower(err.Field()),
			)
			if err.Param() != "" {
				res = strings.ReplaceAll(res, "{param}", err.Param())
			}

			humanReadableError[strings.ToLower(err.Field())] = res
		}

		jsonError, _ := json.Marshal(humanReadableError)

		http.Error(w, string(jsonError), http.StatusUnprocessableEntity)
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
