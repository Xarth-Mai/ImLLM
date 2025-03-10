package utils

import "net/http"

func SetCookies(w http.ResponseWriter, username string, token string) {
	http.SetCookie(w, &http.Cookie{
		Name:     "username",
		Value:    username,
		Path:     "/",
		HttpOnly: true,
	})

	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
	})
}
