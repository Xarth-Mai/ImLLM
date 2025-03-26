package chat

import (
	"encoding/json"
	"github.com/Xarth-Mai/ImLLM/internal/dialog"
	"github.com/Xarth-Mai/ImLLM/internal/utils"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// StreamResponse defines the stream response data structure
type StreamResponse struct {
	ID                string         `json:"id"`
	Object            string         `json:"object"`
	Created           int64          `json:"created"`
	Model             string         `json:"model"`
	SystemFingerprint string         `json:"system_fingerprint"`
	Choices           []StreamChoice `json:"choices"`
}

// StreamChoice defines the choice data structure in stream responses
type StreamChoice struct {
	Index        int            `json:"index"`
	Delta        dialog.Message `json:"delta"`
	Logprobs     any            `json:"logprobs"`
	FinishReason string         `json:"finish_reason"`
}

// HandleChatCompletionsStream handles streaming chat completions API requests
func HandleChatCompletionsStream(w http.ResponseWriter, r *http.Request, req dialog.Request) {
	username := utils.GetUsername(r)

	content, err := dialog.SendRequest(username, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	flusher, ok := w.(http.Flusher)
	if !ok {
		http.Error(w, "Streaming unsupported", http.StatusInternalServerError)
		return
	}

	chunks := strings.Split(content, " ")
	for i, chunk := range chunks {
		streamResponse := StreamResponse{
			ID:                username + strconv.FormatInt(time.Now().Unix(), 10),
			Object:            "chat.completion.chunk",
			Created:           time.Now().Unix(),
			Model:             username,
			SystemFingerprint: username,
			Choices: []StreamChoice{
				{
					Index: 0,
					Delta: dialog.Message{
						Role:    "assistant",
						Content: chunk,
					},
					Logprobs:     nil,
					FinishReason: "",
				},
			},
		}
		if i == len(chunks)-1 {
			streamResponse.Choices[0].FinishReason = "stop"
		}

		b, err := json.Marshal(streamResponse)
		if err != nil {
			continue
		}

		_, err = w.Write([]byte("data: "))
		if err != nil {
			return
		}
		_, err = w.Write(b)
		if err != nil {
			return
		}
		_, err = w.Write([]byte("\n\n"))
		if err != nil {
			return
		}
		flusher.Flush()
	}
}
