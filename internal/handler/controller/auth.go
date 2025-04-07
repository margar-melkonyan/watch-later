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

		helper.SendResponse(w, http.StatusUnprocessableEntity, &helper.Response{
			Messages: humanReadableError,
		})
		return
	}

	err = a.authService.SignUp(&userForm)
	if err != nil {
		helper.SendError(w, http.StatusConflict, helper.MessageResponse{
			Message: err.Error(),
		})
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
		helper.SendResponse(w, http.StatusUnprocessableEntity, &helper.Response{
			Data: humanReadableErrors,
		})
		return
	}

	tokens, err := a.authService.SignIn(&userForm)
	if err != nil {
		helper.SendError(w, http.StatusUnauthorized, helper.MessageResponse{
			Message: err.Error(),
		})
		return
	}

	helper.SendResponse(w, http.StatusOK, &helper.Response{
		Data: tokens,
	})
}

func (a *AuthController) CurrentUser(w http.ResponseWriter, r *http.Request) {
	currentUserEmail, ok := r.Context().Value("user_email").(string)
	if !ok {
		helper.SendError(w, http.StatusConflict, helper.MessageResponse{
			Message: "user not entered",
		})
		return
	}

	user, _ := a.authService.CurrentUser(currentUserEmail)

	helper.SendResponse(w, http.StatusOK, &helper.Response{
		Data: user,
	})
}

func (a *AuthController) RefreshToken(w http.ResponseWriter, r *http.Request) {
	token := r.Header.Get("Authorization")
	if token == "" {
		helper.SendError(w, http.StatusForbidden, helper.MessageResponse{
			Message: "You should be authorized!",
		})
		return
	}

	tokens, err := a.authService.RefreshToken(token)
	if err != nil {
		helper.SendError(w, http.StatusForbidden, helper.MessageResponse{
			Message: err.Error(),
		})
		return
	}

	helper.SendResponse(w, http.StatusOK, &helper.Response{
		Data: tokens,
	})
}
