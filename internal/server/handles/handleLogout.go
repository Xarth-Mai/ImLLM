package handles

import (
	"encoding/json"
	"net/http"
)

// HandleLogout handles the logout API endpoint.
func HandleLogout(w http.ResponseWriter, r *http.Request, userToken *map[string]string) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	username := r.FormValue("username")
	token := r.FormValue("token")

	if username == "" || token == "" {
		http.Error(w, "Missing username or token", http.StatusBadRequest)
		return
	}

	if storedPass, ok := (*userToken)[username]; ok && storedPass == token {
		(*userToken)[username] = ""
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(map[string]string{
			"status":  "success",
			"message": "Logout successful",
		})
		if err != nil {
			http.Error(w, "Encoding response error", http.StatusInternalServerError)
		}
	} else {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
	}
}
