package dialog

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

	// Create a reply channel
	client.ReplyChan = make(chan []byte, 1)
	defer func() { client.ReplyChan = nil }()

	reqBytes, err := json.Marshal(request)
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %v", err)
	}

	if err := client.Conn.WriteMessage(websocket.TextMessage, reqBytes); err != nil {
		return "", fmt.Errorf("failed to send request: %v", err)
	}

	// Wait for reply
	timeout := time.After(30 * time.Second)
	select {
	case reply := <-client.ReplyChan:
		return string(reply), nil
	case <-timeout:
		return "", fmt.Errorf("timeout waiting for reply")
	}
}
