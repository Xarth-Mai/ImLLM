package middleware

import (
	"github.com/Xarth-Mai/ImLLM/internal/utils"
	"net/http"
)

// WebAuth is a middleware that checks if the user is logged in
func WebAuth(next http.HandlerFunc, userToken map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := utils.GetCookieValue(r, "username")
		token := utils.GetCookieValue(r, "token")

		if !utils.ValidateMap(userToken, username, token) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		next(w, r)
	}
}
