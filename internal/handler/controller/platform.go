package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-playground/validator/v10"
	"github.com/margar-melkonyan/watch-later.git/internal/helper"
	"github.com/margar-melkonyan/watch-later.git/internal/repository"
	service "github.com/margar-melkonyan/watch-later.git/internal/service/platforms"
)

type PlatformController struct {
	service *service.PlatformService
}

func NewPlatformController(db *sql.DB) *PlatformController {
	platformRepo := repository.NewPlatformRepository(db)

	return &PlatformController{
		service: service.NewPlatformService(platformRepo),
	}
}

func (controller *PlatformController) GetPlatforms(w http.ResponseWriter, r *http.Request) {
	platforms, err := controller.service.GetPlatforms()
	if err != nil {
		helper.SendError(w, http.StatusInternalServerError, helper.MessageResponse{
			Message: err.Error(),
		})
	}

	helper.SendResponse(w, http.StatusOK, &helper.Response{
		Data: platforms,
	})
}

func (controller *PlatformController) GetPlatform(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(r.PathValue("id"), 0, 64)
	if err != nil {
		helper.SendError(w, http.StatusUnprocessableEntity, helper.MessageResponse{
			Message: "is not valid id",
		})
	}

	platform, err := controller.service.GetPlatform(id)
	if err != nil {
		helper.SendResponse(w, http.StatusNotFound, &helper.Response{
			Data: err.Error(),
		})
		return
	}

	helper.SendResponse(w, http.StatusOK, &helper.Response{
		Data: platform,
	})
}

func (controller *PlatformController) StorePlatform(w http.ResponseWriter, r *http.Request) {
	contentType := w.Header().Get("Content-Type")
	if contentType != "" {
		mediaType := strings.ToLower(strings.TrimSpace(contentType))
		if mediaType != "application/json" {
			helper.SendError(w, http.StatusUnprocessableEntity, helper.MessageResponse{
				Message: "not valid content-type",
			})
			return
		}
	}

	var form repository.Platform
	r.Body = http.MaxBytesReader(w, r.Body, 10<<20)
	err := json.NewDecoder(r.Body).Decode(&form)

	if err != nil {
		helper.SendError(w, http.StatusUnprocessableEntity, helper.MessageResponse{
			Message: err.Error(),
		})
	}

	validate := validator.New()
	err = validate.Struct(form)

	if err != nil {
		errs := err.(validator.ValidationErrors)
		humanReadeable, err := helper.LocalizedValidationMessages(r.Context(), errs)
		if err != nil {
			helper.SendError(w, http.StatusConflict, helper.MessageResponse{
				Message: err.Error(),
			})
		}
		helper.SendResponse(w, http.StatusUnprocessableEntity, &helper.Response{
			Data: humanReadeable,
		})
		return
	}

	if err := controller.service.CreatePlatform(&form); err != nil {
		helper.SendError(w, http.StatusInternalServerError, helper.MessageResponse{
			Message: err.Error(),
		})
	}

	helper.SendResponse(w, http.StatusCreated, nil)
}

func (controller *PlatformController) UpdatePlatform(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(r.PathValue("id"), 0, 64)

	if err != nil {
		helper.SendError(w, http.StatusNotFound, helper.MessageResponse{
			Message: "param is not valid id",
		})
		return
	}

	r.Body = http.MaxBytesReader(w, r.Body, 10<<20)
	var form repository.Platform
	err = json.NewDecoder(r.Body).Decode(&form)
	if err != nil {
		helper.SendError(w, http.StatusConflict, helper.MessageResponse{
			Message: err.Error(),
		})
		return
	}

	validate := validator.New()
	err = validate.Struct(&form)

	if err != nil {
		errs := err.(validator.ValidationErrors)
		humanReadeable, err := helper.LocalizedValidationMessages(r.Context(), errs)
		if err != nil {
			helper.SendError(w, http.StatusConflict, helper.MessageResponse{
				Message: err.Error(),
			})
			return
		}
		helper.SendResponse(w, http.StatusOK, &helper.Response{
			Data: humanReadeable,
		})
		return
	}

	if err := controller.service.UpdatePlatform(&form, id); err != nil {
		helper.SendError(w, http.StatusConflict, helper.MessageResponse{
			Message: "can't update data",
		})
	}
}

func (controller *PlatformController) DeletePlatform(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(r.PathValue("id"), 0, 64)

	if err != nil {
		helper.SendError(w, http.StatusNotFound, helper.MessageResponse{
			Message: "param is not valid id",
		})
		return
	}

	if err := controller.service.DeletePlatform(id); err != nil {
		helper.SendResponse(w, http.StatusNoContent, nil)
	}
}

func (controller *PlatformController) RestorePlatform(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(r.PathValue("id"), 0, 64)

	if err != nil {
		helper.SendError(w, http.StatusNotFound, helper.MessageResponse{
			Message: "param is not valid id",
		})
		return
	}

	if err := controller.service.RestorePlatform(id); err != nil {
		helper.SendError(w, http.StatusNotFound, helper.MessageResponse{
			Message: err.Error(),
		})
		return
	}
}
