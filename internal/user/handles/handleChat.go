package handles

import (
	"github.com/Xarth-Mai/ImLLM/internal/user/utils"
	"net/http"
)

// HandleChat is the main function for the chat API
func HandleChat() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username := utils.GetCookieValue(r, "username")
		if username == "" {
			return
		}
	}
}
