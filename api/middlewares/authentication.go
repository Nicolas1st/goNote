package middlewares

import (
	"net/http"

	"github.com/Nicolas1st/goNote/api/auth"
)

func AuthenticationMiddleware(next http.Handler, handlerIfNotLoggedIn http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if auth.IsAuthenticated(w, r) {
			next.ServeHTTP(w, r)
		} else {
			handlerIfNotLoggedIn(w, r)
		}
	})
}
