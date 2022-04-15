package auth

import "net/http"

func login(w http.ResponseWriter, r *http.Request) {
	validCookie := &http.Cookie{
		Name:     validCookieFields.Name,
		Value:    validCookieFields.Value,
		Path:     validCookieFields.Path,
		HttpOnly: validCookieFields.HttpOnly,
	}

	http.SetCookie(w, validCookie)
}

func logout(w http.ResponseWriter, r *http.Request) {
	expirationCookie := &http.Cookie{
		Name:     validCookieFields.Name,
		Value:    validCookieFields.Value,
		Path:     validCookieFields.Path,
		HttpOnly: validCookieFields.HttpOnly,
		MaxAge:   validCookieFields.MaxAge,
	}

	http.SetCookie(w, expirationCookie)
}

func IsAuthenticated(w http.ResponseWriter, r *http.Request) bool {
	// not implemented properly yet
	authCookie, err := r.Cookie(validCookieFields.Name)
	if err != nil {
		return false
	}

	if authCookie.Name == validCookieFields.Name && authCookie.Value == validCookieFields.Value {
		return true
	}

	return false
}
