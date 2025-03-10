package openai

import (
	"github.com/Xarth-Mai/ImLLM/internal/server/api/openai/chat"
	"github.com/Xarth-Mai/ImLLM/internal/server/api/openai/models"
	"net/http"
	"strings"
)

// HandleOpenAI is the main function for the OpenAI API
func HandleOpenAI(w http.ResponseWriter, r *http.Request) {
	username, ok := r.Context().Value("username").(string)
	if !ok || username == "" {
		http.Error(w, "username missing", http.StatusUnauthorized)
		return
	}
	apiPath := strings.TrimPrefix(r.URL.Path, "/openai/")

	switch apiPath {
	case "models":
		models.HandleModels(w, username)
	case "chat/completions":
		chat.HandleChatCompletions(w, r, username)
	default:
		http.Error(w, "Not Found", http.StatusNotFound)
	}
}
