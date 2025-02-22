package router

import (
	"database/sql"
	"github.com/margar-melkonyan/watch-later.git/internal/repository"
	"github.com/margar-melkonyan/watch-later.git/internal/service/notifications"
	"net/http"
)

func notificationsRoutes(db *sql.DB) *http.ServeMux {
	notifications := http.NewServeMux()
	notificationRepo := repository.NewNotificationRepository(db)
	_ = service.NewNotificationService(notificationRepo)

	notifications.HandleFunc("GET /unread", func(w http.ResponseWriter, r *http.Request) {})
	notifications.HandleFunc("GET /ws", func(w http.ResponseWriter, r *http.Request) {})
	notifications.HandleFunc("POST /{id}/mark-as-read", func(w http.ResponseWriter, r *http.Request) {})
	notifications.HandleFunc("POST /{id}/multiple-mark-as-read", func(w http.ResponseWriter, r *http.Request) {})

	return notifications
}
