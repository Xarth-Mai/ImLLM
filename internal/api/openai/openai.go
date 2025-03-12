package openai

import (
	"github.com/Xarth-Mai/ImLLM/internal/api/openai/chat"
	"github.com/Xarth-Mai/ImLLM/internal/api/openai/models"
	"github.com/Xarth-Mai/ImLLM/internal/middleware"
	"net/http"
	"strings"
)

// HandleOpenAI is the main function for the OpenAI API
func HandleOpenAI(userPasswd map[string]string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		apiPath := strings.TrimPrefix(r.URL.Path, "/openai/")

		switch apiPath {
		case "models":
			middleware.ApiAuth(models.HandleModels, userPasswd)(w, r)
		case "chat/completions":
			middleware.ApiAuth(chat.HandleChatCompletions, userPasswd)(w, r)
		default:
			http.Error(w, "Not Found", http.StatusNotFound)
		}
	}
}
