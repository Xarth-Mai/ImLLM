package utils

import "net/http"

// GetUsername returns the username from the request context
func GetUsername(r *http.Request) string {
	username := r.Context().Value("username").(string)
	return username
}
