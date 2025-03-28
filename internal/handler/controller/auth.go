package controller

import (
	"database/sql"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/margar-melkonyan/watch-later.git/internal/common"
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
	if contentType == "" {
		helper.SendError(w, http.StatusInternalServerError, helper.MessageResponse{
			Message: "content-type is required",
		})
		return
	}

	if contentType != "" &&
		strings.ToLower(strings.TrimSpace(contentType)) != "application/json" {
		helper.SendError(w, http.StatusBadRequest, helper.MessageResponse{
			Message: "not valid content-type",
		})
		return
	}
	r.Body = http.MaxBytesReader(w, r.Body, 1024*1024*10)
	defer r.Body.Close()

	var userForm repository.User

	if err := json.NewDecoder(r.Body).Decode(&userForm); err != nil {
		if err == io.EOF {
			helper.SendError(w, http.StatusBadRequest, helper.MessageResponse{
				Message: "not valid json",
			})
			return
		}

		helper.SendError(w, http.StatusBadRequest, helper.MessageResponse{
			Message: err.Error(),
		})
		return
	}

	validate := validator.New()
	err := validate.Struct(userForm)

	if err != nil {
		errs := err.(validator.ValidationErrors)

		humanReadableError, err := helper.LocalizedValidationMessages(r.Context(), errs)
		if err != nil {
			helper.SendError(w, http.StatusInternalServerError, helper.MessageResponse{
				Message: err.Error(),
			})
			return
		}

		helper.SendResponse(w, http.StatusUnprocessableEntity, helper.Response{
			Messages: humanReadableError,
		})
		return
	}

	ok := a.authService.SignUp(&userForm)

	if ok != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (a *AuthController) SignIn(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	if contentType == "" {
		helper.SendError(w, http.StatusInternalServerError, helper.MessageResponse{
			Message: "content type is required",
		})
	}

	if contentType != "" &&
		strings.TrimSpace(contentType) != "application/json" {
		helper.SendError(w, http.StatusInternalServerError, helper.MessageResponse{
			Message: "not valid content-type",
		})
	}
	r.Body = http.MaxBytesReader(w, r.Body, 10*1024*1024)
	defer r.Body.Close()

	var userForm common.SignInUser
	if err := json.NewDecoder(r.Body).Decode(&userForm); err != nil {
		helper.SendError(w, http.StatusBadRequest, helper.MessageResponse{
			Message: "not valid json",
		})
		return
	}

	validate := validator.New()
	err := validate.Struct(userForm)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		humanReadableErrors, err := helper.LocalizedValidationMessages(r.Context(), errs)
		if err != nil {
			helper.SendError(w, http.StatusInternalServerError, helper.MessageResponse{
				Message: err.Error(),
			})
			return
		}
		helper.SendResponse(w, http.StatusUnprocessableEntity, helper.Response{
			Data: humanReadableErrors,
		})
	}

	_, err = a.authService.SignIn(&userForm)

	if err != nil {
		helper.SendError(w, http.StatusUnauthorized, helper.MessageResponse{
			Message: err.Error(),
		})
		return
	}

}

func (a *AuthController) SignOut(w http.ResponseWriter, r *http.Request) {

}
