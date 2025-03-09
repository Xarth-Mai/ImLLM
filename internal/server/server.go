package server

import (
	"net/http"
	"strings"
)

func Run(userPasswd map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiPath := strings.TrimPrefix(r.URL.Path, "/api/")

		switch apiPath {
		case "login":
			handleLogin(w, r, userPasswd)
		default:
			http.Error(w, "Unknown API endpoint", http.StatusNotFound)
		}
	}
}
