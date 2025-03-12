package chat

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"time"
)

// Request defines the chat request data structure
type Request struct {
	Model    string    `json:"model"`
	Messages []Message `json:"messages"`
}

// Message defines the message data structure
type Message struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

// SendRequest sends a chat request to the specified user
func SendRequest(username string, request Request) (string, error) {
	clientsMu.RLock()
	client, online := clients[username]
	clientsMu.RUnlock()

	if !online {
		return "", fmt.Errorf("user %s is not online", username)
	}

	client.Mutex.Lock()
	defer client.Mutex.Unlock()

	reqBytes, err := json.Marshal(request)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %v", err)
	}

	if err := client.Conn.WriteMessage(websocket.TextMessage, reqBytes); err != nil {
		return "", fmt.Errorf("failed to send request: %v", err)
	}

	err = client.Conn.SetReadDeadline(time.Now().Add(600 * time.Second))
	if err != nil {
		return "", err
	}

	_, reply, err := client.Conn.ReadMessage()
	if err != nil {
		return "", fmt.Errorf("failed to read reply: %v", err)
	}

	return string(reply), nil
}
