package middleware

import "net/http"

type Middleware func(http.Handler) http.Handler

func Stack(middlwares ...Middleware) Middleware {
	return func(next http.Handler) http.Handler {
		for i := len(middlwares) - 1; i >= 0; i-- {
			middlware := middlwares[i]
			next = middlware(next)
		}

		return next
	}
}
