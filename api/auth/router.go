package auth

import "net/http"

func NewAuthRouter() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/auth/login", login)
	mux.HandleFunc("/auth/logout", logout)

	return mux
}
