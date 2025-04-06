package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/margar-melkonyan/watch-later.git/internal/helper"
	service "github.com/margar-melkonyan/watch-later.git/internal/service/users"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.URL.Path, "auth") || strings.Contains(r.URL.Path, "current-user") {
			token := r.Header.Get("Authorization")
			if token == "" {
				helper.SendError(w, http.StatusForbidden, helper.MessageResponse{
					Message: "You should be authorized!",
				})
				return
			}

			claims, err := service.CheckTokenIsNotExpired(token, "JWT_ACCESS_TOKEN")
			if err != nil {
				helper.SendError(w, http.StatusForbidden, helper.MessageResponse{
					Message: err.Error(),
				})
				return
			}

			req := context.WithValue(r.Context(), "user_email", claims.Sub.Email)
			r = r.WithContext(req)
		}

		next.ServeHTTP(w, r)
	})
}
