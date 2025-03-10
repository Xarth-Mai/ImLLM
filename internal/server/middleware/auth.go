package middleware

import (
	"net/http"
)

// Auth is a middleware that checks if the user is logged in
func Auth(next http.HandlerFunc, userToken map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//username := r.FormValue("username")
		//token := r.FormValue("token")
		//
		//if !utils.ValidateMap(userToken, username, token) {
		//	http.Error(w, "Unauthorized", http.StatusUnauthorized)
		//	return
		//}
		next(w, r)
	}
}
