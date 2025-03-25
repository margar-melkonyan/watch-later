package middleware

import (
	"context"
	"net/http"
	"strings"
)

func SetLocale(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		locale := "en"

		if lang := r.Header.Get("Accept-Language"); lang != "" {
			locale = strings.ToLower(strings.Split(lang, "-")[0])
		}

		ctx := context.WithValue(r.Context(), "locale", locale)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
