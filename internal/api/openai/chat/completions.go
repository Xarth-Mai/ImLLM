package chat

import (
	"encoding/json"
	"github.com/Xarth-Mai/ImLLM/internal/utils"
	"net/http"
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

// Choice defines the choice data structure
type Choice struct {
	Index        int     `json:"index"`
	Message      Message `json:"message"`
	Logprobs     any     `json:"logprobs"`
	FinishReason string  `json:"finish_reason"`
}

// Message defines the message data structure
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
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

// HandleChatCompletions handles the chat completions API request
func HandleChatCompletions(w http.ResponseWriter, r *http.Request) {
	username := utils.GetUsername(r)

	response := Response{
		ID:                username,
		Object:            "chat.completion",
		Created:           1677652288,
		Model:             "gpt-4o-mini",
		SystemFingerprint: "fp_44709d6fcb",
		Choices: []Choice{
			{
				Index: 0,
				Message: Message{
					Role:    "assistant",
					Content: "\n\nHello there, how may I assist you today?",
				},
				Logprobs:     nil,
				FinishReason: "stop",
			},
		},
		Usage: Usage{
			PromptTokens:     9,
			CompletionTokens: 12,
			TotalTokens:      21,
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
