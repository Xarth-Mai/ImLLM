package handles

import (
	"encoding/json"
	"github.com/Xarth-Mai/ImLLM/internal/server/utils"
	"log"
	"net/http"
)

// HandleLogin handles the login API endpoint.
func HandleLogin(w http.ResponseWriter, r *http.Request, userPasswd *map[string]string, userToken *map[string]string) {
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

	storedPass := utils.ValidateMap(*userPasswd, username, password)

	if storedPass {
		token := utils.GenerateToken()

		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(map[string]string{
			"status":  "success",
			"message": "Login successful",
			"token":   token,
		})
		if err != nil {
			http.Error(w, "Encoding response error", http.StatusInternalServerError)
			return
		}
		(*userToken)[username] = token

		log.Printf("User '%s' logged in", username)
	} else {
		http.Error(w, "Invalid username or password", http.StatusUnauthorized)
	}
}
