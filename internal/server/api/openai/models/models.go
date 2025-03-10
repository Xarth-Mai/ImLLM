package models

import (
	"encoding/json"
	"net/http"
)

// Model defines the model structure for the Models API
type Model struct {
	ID      string `json:"id"`
	Object  string `json:"object"`
	Created int64  `json:"created"`
	OwnedBy string `json:"owned_by"`
}

// Response defines the response structure for the Models API
type Response struct {
	Object string  `json:"object"`
	Data   []Model `json:"data"`
}

// HandleModels is the main function for the Models API
func HandleModels(w http.ResponseWriter, username string) {
	w.Header().Set("Content-Type", "application/json")

	models := []Model{
		{
			ID:      username,
			Object:  "model",
			Created: 1700000000,
			OwnedBy: username,
		},
	}

	response := Response{
		Object: "list",
		Data:   models,
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
