package auth

import "net/http"

var validCookieFields *http.Cookie = &http.Cookie{
	Name:     "auth_cookie",
	Value:    "auth_cookie_value",
	HttpOnly: true,
	Path:     "/",
	MaxAge:   -1,
}
