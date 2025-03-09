package openai

import (
	"github.com/Xarth-Mai/ImLLM/internal/server/api/openai/internal/chat"
	"github.com/Xarth-Mai/ImLLM/internal/server/api/openai/internal/models"
	"net/http"
)

// HandleOpenAI is the main function for the OpenAI API
func HandleOpenAI(w http.ResponseWriter, r *http.Request, userToken *map[string]string) {
	switch r.URL.Path {
	case "/models":
		models.HandleModels(w, r, userToken)
	case "/chat/completions":
		chat.HandleChatCompletions(w, r, userToken)
	default:
		http.Error(w, "Not Found", http.StatusNotFound)
	}
}
