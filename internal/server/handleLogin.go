package server

import (
	"encoding/json"
	"net/http"
)

// handleLogin handles the login API endpoint.
func handleLogin(w http.ResponseWriter, r *http.Request, userPasswd map[string]string) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	if err := r.ParseForm(); err != nil {
		http.Error(w, "Error parsing form data", http.StatusBadRequest)
		return
	}

	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "" || password == "" {
		http.Error(w, "Missing username or password", http.StatusBadRequest)
		return
	}

	if storedPass, ok := userPasswd[username]; ok && storedPass == password {
		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(map[string]string{
			"status":  "success",
			"message": "Login successful",
			"token":   "token...",
		})
		if err != nil {
			http.Error(w, "Encoding response error", http.StatusInternalServerError)
		}
	} else {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
	}
}
