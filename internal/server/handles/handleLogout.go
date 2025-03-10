package handles

import (
	"github.com/Xarth-Mai/ImLLM/internal/server/utils"
	"log"
	"net/http"
)

// HandleLogout handles logout requests
func HandleLogout(userToken *map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := utils.GetCookieValue(r, "username")

		(*userToken)[username] = ""

		utils.SetCookies(w, "", "")

		log.Printf("User '%s' logged out", username)
		http.Redirect(w, r, "/login.html", http.StatusFound)
	}
}
