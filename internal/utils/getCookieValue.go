package utils

import "net/http"

func GetCookieValue(r *http.Request, key string) string {
	value, err := r.Cookie(key)
	if err != nil {
		return ""
	}
	return value.Value
}
