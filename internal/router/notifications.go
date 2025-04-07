package router

import (
	"database/sql"
	"net/http"

	"github.com/margar-melkonyan/watch-later.git/internal/handler/controller"
)

func notificationsRoutes(db *sql.DB) *http.ServeMux {
	notifications := http.NewServeMux()
	notificationController := controller.NewNotificationController(db)

	notifications.HandleFunc("GET /unread", notificationController.Unread)
	notifications.HandleFunc("POST /{id}/mark-as-read", notificationController.MarkAsRead)
	notifications.HandleFunc("POST /multiple-mark-as-read", notificationController.MultipleMarkAsRead)

	return notifications
}
