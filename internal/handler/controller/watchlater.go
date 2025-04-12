package controller

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-playground/validator/v10"
	"github.com/margar-melkonyan/watch-later.git/internal/helper"
	"github.com/margar-melkonyan/watch-later.git/internal/repository"
	service "github.com/margar-melkonyan/watch-later.git/internal/service/watchlaters"
)

type WatchLaterController struct {
	service *service.WatchLaterService
}

func NewWatchLaterController(db *sql.DB) *WatchLaterController {
	repo := repository.NewWatchLaterRepository(db)
	return &WatchLaterController{
		service: service.NewWatchLaterService(repo),
	}
}

func (c *WatchLaterController) GetWatchLater(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(r.PathValue("id"), 0, 64)
	if err != nil {
		helper.SendError(w, http.StatusBadRequest, helper.MessageResponse{
			Message: "is not valid id",
		})
	}

	watchLater, err := c.service.GetWatchLater(id)
	if err != nil {
		helper.SendResponse(w, http.StatusNotFound, &helper.Response{
			Data: []struct{}{},
		})
		return
	}

	helper.SendResponse(w, http.StatusOK, &helper.Response{
		Data: watchLater,
	})
}

func (c *WatchLaterController) GetWatchLaters(w http.ResponseWriter, r *http.Request) {
	watchLaters, err := c.service.GetWatchLaters()
	if err != nil {
		helper.SendError(
			w, http.StatusNotFound,
			helper.MessageResponse{
				Message: err.Error(),
			},
		)
		return
	}

	helper.SendResponse(w, http.StatusOK, &helper.Response{
		Data: watchLaters,
	})
}

func (c *WatchLaterController) GetWatchLatersByCategory(w http.ResponseWriter, r *http.Request) {
	cateogyrID, err := strconv.ParseUint(r.PathValue("cateogry_id"), 0, 64)
	if err != nil {
		helper.SendError(w, http.StatusConflict, helper.MessageResponse{
			Message: "is not correct category id",
		})
		return
	}

	watchLaters, err := c.service.GetWatchLatersByCategory(cateogyrID)
	if err != nil {
		helper.SendError(
			w, http.StatusNotFound,
			helper.MessageResponse{
				Message: err.Error(),
			},
		)
		return
	}

	helper.SendResponse(w, http.StatusOK, &helper.Response{
		Data: watchLaters,
	})
}

func (c *WatchLaterController) GetWatchLatersByPlatform(w http.ResponseWriter, r *http.Request) {
	platformID, err := strconv.ParseUint(r.PathValue("platform_id"), 0, 64)
	if err != nil {
		helper.SendError(w, http.StatusConflict, helper.MessageResponse{
			Message: "is not correct category id",
		})
		return
	}

	watchLaters, err := c.service.GetWatchLatersByPlatform(platformID)
	if err != nil {
		helper.SendError(
			w, http.StatusNotFound,
			helper.MessageResponse{
				Message: err.Error(),
			},
		)
		return
	}

	helper.SendResponse(w, http.StatusOK, &helper.Response{
		Data: watchLaters,
	})
}

func (c *WatchLaterController) StoreWatchLater(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	if contentType != "" {
		if contentType != "application/json" {
			helper.SendError(w, http.StatusUnsupportedMediaType, helper.MessageResponse{
				Message: "not supported content-type",
			})
			return
		}
	}

	var watchLater repository.WatchLater
	r.Body = http.MaxBytesReader(w, r.Body, 10<<20)
	err := json.NewDecoder(r.Body).Decode(&watchLater)

	if err != nil {
		helper.SendError(w, http.StatusBadRequest, helper.MessageResponse{
			Message: err.Error(),
		})
		return
	}

	validate := validator.New()
	err = validate.Struct(&watchLater)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		humanReadable, err := helper.LocalizedValidationMessages(r.Context(), errs)
		if err != nil {
			helper.SendError(w, http.StatusInternalServerError, helper.MessageResponse{
				Message: err.Error(),
			})
			return
		}
		helper.SendResponse(w, http.StatusUnprocessableEntity, &helper.Response{
			Data: humanReadable,
		})
		return
	}

	if err := c.service.StoreWatchLater(&watchLater); err != nil {
		helper.SendError(w, http.StatusInternalServerError, helper.MessageResponse{
			Message: err.Error(),
		})
		return
	}
}

func (c *WatchLaterController) UpdateWatchLater(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(r.PathValue("id"), 0, 64)
	if err != nil {
		helper.SendError(w, http.StatusConflict, helper.MessageResponse{
			Message: "is not valid id",
		})
		return
	}

	contentType := r.Header.Get("Content-Type")
	if contentType != "" {
		if contentType != "application/json" {
			helper.SendError(w, http.StatusUnsupportedMediaType, helper.MessageResponse{
				Message: "not supported content-type",
			})
			return
		}
	}

	var watchLater repository.WatchLater
	r.Body = http.MaxBytesReader(w, r.Body, 10<<20)
	err = json.NewDecoder(r.Body).Decode(&watchLater)

	if err != nil {
		helper.SendError(w, http.StatusBadRequest, helper.MessageResponse{
			Message: err.Error(),
		})
		return
	}

	validate := validator.New()
	err = validate.Struct(&watchLater)
	if err != nil {
		errs := err.(validator.ValidationErrors)
		humanReadable, err := helper.LocalizedValidationMessages(r.Context(), errs)
		if err != nil {
			helper.SendError(w, http.StatusInternalServerError, helper.MessageResponse{
				Message: err.Error(),
			})
			return
		}
		helper.SendResponse(w, http.StatusUnprocessableEntity, &helper.Response{
			Data: humanReadable,
		})
		return
	}

	if err := c.service.UpdateWatchLater(&watchLater, id); err != nil {
		helper.SendError(w, http.StatusInternalServerError, helper.MessageResponse{
			Message: err.Error(),
		})
		return
	}
}

func (c *WatchLaterController) DeleteWatchLater(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(r.PathValue("id"), 0, 64)
	if err != nil {
		helper.SendError(w, http.StatusConflict, helper.MessageResponse{
			Message: "is not valid id",
		})
		return
	}

	if err := c.service.DeleteWatchLater(id); err != nil {
		helper.SendError(w, http.StatusNoContent, helper.MessageResponse{
			Message: err.Error(),
		})
	}
}
