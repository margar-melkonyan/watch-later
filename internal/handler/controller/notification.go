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
	service "github.com/margar-melkonyan/watch-later.git/internal/service/notifications"
)

type NotificationController struct {
	service *service.NotificationService
}

func NewNotificationController(db *sql.DB) *NotificationController {
	notificationRepo := repository.NewNotificationRepository(db)
	userRepository := repository.NewUserRepository(db)

	return &NotificationController{
		service: service.NewNotificationService(
			notificationRepo,
			userRepository,
		),
	}
}

func (controller *NotificationController) Unread(w http.ResponseWriter, r *http.Request) {
	notifications, err := controller.service.GetUnreadNotifications(r.Context())
	if err != nil {
		helper.SendError(w, http.StatusInternalServerError, helper.MessageResponse{
			Message: "Error getting unread notifications",
		})
		return
	}

	helper.SendResponse(w, http.StatusOK, &helper.Response{
		Data: notifications,
	})
}

func (controller *NotificationController) MarkAsRead(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.ParseUint(r.PathValue("id"), 0, 64)
	if err != nil {
		helper.SendError(w, http.StatusUnprocessableEntity, helper.MessageResponse{
			Message: "Invalid notification id",
		})
	}

	if err := controller.service.MarkAsReadNotification(id); err != nil {
		helper.SendError(w, http.StatusInternalServerError, helper.MessageResponse{
			Message: "Error marking notification as read",
		})
	}

	helper.SendResponse(w, http.StatusOK, nil)
}

func (controller *NotificationController) MultipleMarkAsRead(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	if contentType != "" {
		mediaType := strings.ToLower(strings.TrimSpace(contentType))
		if mediaType != "application/json" {
			helper.SendError(w, http.StatusUnsupportedMediaType, helper.MessageResponse{
				Message: "not valid content-type",
			})
			return
		}
	}

	type data struct {
		Ids []uint64 `json:"ids" validate:"required,min=1"`
	}

	ids := &data{
		Ids: make([]uint64, 0),
	}

	r.Body = http.MaxBytesReader(w, r.Body, 10<<20)
	err := json.NewDecoder(r.Body).Decode(ids)
	if err != nil {
		helper.SendError(w, http.StatusUnprocessableEntity, helper.MessageResponse{
			Message: err.Error(),
		})
		return
	}

	validate := validator.New()
	err = validate.Struct(ids)

	if err != nil {
		errs := err.(validator.ValidationErrors)
		humanReadable, err := helper.LocalizedValidationMessages(r.Context(), errs)
		if err != nil {
			helper.SendError(w, http.StatusUnprocessableEntity, helper.MessageResponse{
				Message: err.Error(),
			})
			return
		}
		helper.SendResponse(w, http.StatusUnprocessableEntity, &helper.Response{
			Data: humanReadable,
		})
		return
	}

	controller.service.MultipleMarkAsRead(ids.Ids)
}
