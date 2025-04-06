package middleware

import (
	"net/http"
	"strings"

	"github.com/margar-melkonyan/watch-later.git/internal/helper"
	service "github.com/margar-melkonyan/watch-later.git/internal/service/users"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.URL.Path, "auth") {
			token := r.Header.Get("Authorization")

			if token == "" {
				helper.SendError(w, http.StatusForbidden, helper.MessageResponse{
					Message: "You should be authorized!",
				})
				return
			}

			_, err := service.CheckTokenIsNotExpired(token, "JWT_ACCESS_TOKEN")
			if err != nil {
				helper.SendError(w, http.StatusForbidden, helper.MessageResponse{
					Message: err.Error(),
				})
				return
			}
		}

		next.ServeHTTP(w, r)
	})
}
