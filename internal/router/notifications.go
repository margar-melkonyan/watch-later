package router

import (
	"database/sql"
	"net/http"
)

func notificationsRoutes(db *sql.DB) *http.ServeMux {
	notifications := http.NewServeMux()

	notifications.HandleFunc("GET /unread", func(w http.ResponseWriter, r *http.Request) {})
	notifications.HandleFunc("GET /ws", func(w http.ResponseWriter, r *http.Request) {})
	notifications.HandleFunc("POST /{id}/mark-as-read", func(w http.ResponseWriter, r *http.Request) {})
	notifications.HandleFunc("POST /{id}/multiple-mark-as-read", func(w http.ResponseWriter, r *http.Request) {})

	return notifications
}
