package openai

import (
	"github.com/Xarth-Mai/ImLLM/internal/server/api/openai/internal/chat"
	"github.com/Xarth-Mai/ImLLM/internal/server/api/openai/internal/models"
	"net/http"
)

// HandleOpenAI is the main function for the OpenAI API
func HandleOpenAI() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.URL.Path {
		case "/models":
			models.HandleModels(w, r)
		case "/chat/completions":
			chat.HandleChatCompletions(w, r)
		default:
			http.Error(w, "Not Found", http.StatusNotFound)
		}
	}
}
