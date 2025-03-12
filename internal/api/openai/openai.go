package openai

import (
	"github.com/Xarth-Mai/ImLLM/internal/api/openai/chat"
	"github.com/Xarth-Mai/ImLLM/internal/api/openai/models"
	"net/http"
	"strings"
)

// HandleOpenAI is the main function for the OpenAI API
func HandleOpenAI(w http.ResponseWriter, r *http.Request) {
	apiPath := strings.TrimPrefix(r.URL.Path, "/openai/")

	switch apiPath {
	case "models":
		models.HandleModels(w, r)
	case "chat/completions":
		chat.HandleChatCompletions(w, r)
	default:
		http.Error(w, "Not Found", http.StatusNotFound)
	}

}
