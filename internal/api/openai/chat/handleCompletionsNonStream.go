package chat

import (
	"encoding/json"
	"github.com/Xarth-Mai/ImLLM/internal/dialog"
	"github.com/Xarth-Mai/ImLLM/internal/utils"
	"net/http"
	"strconv"
	"time"
)

// Response defines the response data structure
type Response struct {
	ID                string   `json:"id"`
	Object            string   `json:"object"`
	Created           int64    `json:"created"`
	Model             string   `json:"model"`
	SystemFingerprint string   `json:"system_fingerprint"`
	Choices           []Choice `json:"choices"`
	Usage             Usage    `json:"usage"`
}

// Choice defines the choice data structure in non stream responses
type Choice struct {
	Index        int            `json:"index"`
	Message      dialog.Message `json:"message"`
	Logprobs     any            `json:"logprobs"`
	FinishReason string         `json:"finish_reason"`
}

// Usage defines the usage data structure
type Usage struct {
	PromptTokens            int                     `json:"prompt_tokens"`
	CompletionTokens        int                     `json:"completion_tokens"`
	TotalTokens             int                     `json:"total_tokens"`
	CompletionTokensDetails CompletionTokensDetails `json:"completion_tokens_details"`
}

// CompletionTokensDetails defines the completion tokens details
type CompletionTokensDetails struct {
	ReasoningTokens          int `json:"reasoning_tokens"`
	AcceptedPredictionTokens int `json:"accepted_prediction_tokens"`
	RejectedPredictionTokens int `json:"rejected_prediction_tokens"`
}

// HandleChatCompletionsNonStream handles non-streaming chat completions API requests
func HandleChatCompletionsNonStream(w http.ResponseWriter, r *http.Request, req dialog.Request) {
	username := utils.GetUsername(r)

	content, err := dialog.SendRequest(username, req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response := Response{
		ID:                username + strconv.FormatInt(time.Now().Unix(), 10),
		Object:            "chat.completion",
		Created:           time.Now().Unix(),
		Model:             username,
		SystemFingerprint: username,
		Choices: []Choice{
			{
				Index: 0,
				Message: dialog.Message{
					Role:    "assistant",
					Content: content,
				},
				Logprobs:     nil,
				FinishReason: "stop",
			},
		},
		Usage: Usage{
			PromptTokens:     1000,
			CompletionTokens: 1000,
			TotalTokens:      2000,
			CompletionTokensDetails: CompletionTokensDetails{
				ReasoningTokens:          0,
				AcceptedPredictionTokens: 0,
				RejectedPredictionTokens: 0,
			},
		},
	}

	w.Header().Set("Content-Type", "application/json")

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
