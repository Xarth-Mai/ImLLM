package handles

import (
	"encoding/json"
	"net/http"
)

// HandleLogout handles logout requests
func HandleLogout(userToken *map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			return
		}

		if err := r.ParseForm(); err != nil {
			http.Error(w, "Error parsing form data", http.StatusBadRequest)
			return
		}

		username := r.FormValue("username")

		(*userToken)[username] = ""
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(map[string]string{
			"status":  "success",
			"message": "Logout successful",
		}); err != nil {
			http.Error(w, "Encoding response error", http.StatusInternalServerError)
		}
	}
}
