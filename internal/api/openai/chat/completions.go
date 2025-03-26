package chat

import (
	"encoding/json"
	"github.com/Xarth-Mai/ImLLM/internal/dialog"
	"net/http"
)

// reference: https://platform.openai.com/docs/api-reference/chat

// HandleChatCompletions handles chat completions API requests
func HandleChatCompletions(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	var req dialog.Request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if req.Stream == true {
		HandleChatCompletionsStream(w, r, req)
	} else {
		HandleChatCompletionsNonStream(w, r, req)
	}
}
