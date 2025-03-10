package openai

import (
	"github.com/Xarth-Mai/ImLLM/internal/server/api/openai/internal/chat"
	"github.com/Xarth-Mai/ImLLM/internal/server/api/openai/internal/models"
	"net/http"
)

// HandleOpenAI is the main function for the OpenAI API
func HandleOpenAI(w http.ResponseWriter, r *http.Request) {
	username, ok := r.Context().Value("username").(string)
	if !ok || username == "" {
		http.Error(w, "username missing", http.StatusUnauthorized)
		return
	}

	switch r.URL.Path {
	case "/openai/models":
		models.HandleModels(w, username)
	case "/openai/chat/completions":
		chat.HandleChatCompletions(w, r, username)
	default:
		http.Error(w, "Not Found", http.StatusNotFound)
	}
}
