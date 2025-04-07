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
	service "github.com/margar-melkonyan/watch-later.git/internal/service/categories"
)

type CategoryController struct {
	service *service.CategoryService
}

func NewCategoryController(db *sql.DB) *CategoryController {
	categoryRepo := repository.NewCategoryRepository(db)
	return &CategoryController{
		service: service.NewCategoryService(categoryRepo),
	}
}

func (controller *CategoryController) GetCategories(w http.ResponseWriter, r *http.Request) {
	categories, err := controller.service.GetCategories()
	if err != nil {
		helper.SendError(w, http.StatusNotFound, helper.MessageResponse{})
		return
	}

	helper.SendResponse(w, http.StatusOK, &helper.Response{
		Data: categories,
	})
}

func (controller *CategoryController) GetCategoryById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(r.PathValue("id"), 0, 64)
	if err != nil {
		helper.SendError(
			w, http.StatusInternalServerError, helper.MessageResponse{
				Message: "is not valid id",
			},
		)
		return
	}

	category, err := controller.service.GetCategory(id)

	if category == nil {
		helper.SendError(
			w, http.StatusNotFound, helper.MessageResponse{
				Message: err.Error(),
			},
		)
		return
	}

	helper.SendResponse(
		w, http.StatusOK, &helper.Response{
			Data: category,
		},
	)
}

func (controller *CategoryController) StoreCategory(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	if contentType != "" {
		mediaType := strings.ToLower(strings.TrimSpace(contentType))
		if mediaType != "application/json" {
			w.WriteHeader(http.StatusUnprocessableEntity)
			return
		}
	}

	var form repository.Category
	r.Body = http.MaxBytesReader(w, r.Body, 10<<20)
	err := json.NewDecoder(r.Body).Decode(&form)

	if err != nil {
		helper.SendError(
			w, http.StatusInternalServerError, helper.MessageResponse{
				Message: err.Error(),
			},
		)
		return
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
			return
		}
		helper.SendResponse(w, http.StatusUnprocessableEntity, &helper.Response{
			Data: humanReadeable,
		})
		return
	}

	if err := controller.service.CreateCategory(&form); err != nil {
		helper.SendError(
			w, http.StatusInternalServerError, helper.MessageResponse{
				Message: err.Error(),
			},
		)
		return
	}

	helper.SendResponse(
		w, http.StatusCreated, nil,
	)
}

func (controller *CategoryController) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(r.PathValue("id"), 0, 64)
	if err != nil {
		helper.SendError(
			w, http.StatusNotFound, helper.MessageResponse{
				Message: "param is not valid id",
			},
		)
		return
	}

	contentType := r.Header.Get("Content-Type")
	if contentType != "" {
		mediaType := strings.ToLower(strings.TrimSpace(contentType))
		if mediaType != "application/json" {
			w.WriteHeader(http.StatusUnprocessableEntity)
			helper.SendError(w, http.StatusUnprocessableEntity, helper.MessageResponse{
				Message: "not valid body",
			})
			return
		}
	}

	var form repository.Category
	r.Body = http.MaxBytesReader(w, r.Body, 10<<20)
	err = json.NewDecoder(r.Body).Decode(&form)

	if err != nil {
		helper.SendError(w, http.StatusInternalServerError, helper.MessageResponse{
			Message: err.Error(),
		})
		return
	}

	if err := controller.service.UpdateCategory(&form, id); err != nil {
		helper.SendError(w, http.StatusNotFound, helper.MessageResponse{
			Message: err.Error(),
		})
		return
	}
}

func (controller *CategoryController) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(r.PathValue("id"), 0, 64)

	if err != nil {
		helper.SendError(w, http.StatusNotFound, helper.MessageResponse{
			Message: "param is not valid id",
		})
		return
	}

	if err := controller.service.DeleteCategory(id); err != nil {
		helper.SendError(w, http.StatusNotFound, helper.MessageResponse{
			Message: err.Error(),
		})
		return
	}
}

func (controller *CategoryController) RestoreCategory(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(r.PathValue("id"), 0, 64)

	if err != nil {
		helper.SendError(w, http.StatusNotFound, helper.MessageResponse{
			Message: "param is not valid id",
		})
		return
	}

	if err := controller.service.RestoreCategory(id); err != nil {
		helper.SendError(w, http.StatusNotFound, helper.MessageResponse{
			Message: err.Error(),
		})
		return
	}
}
