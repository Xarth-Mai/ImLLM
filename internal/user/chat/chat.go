package chat

import (
	"fmt"
	"github.com/Xarth-Mai/ImLLM/internal/utils"
	"github.com/gorilla/websocket"
	"net/http"
	"sync"
)

// upgrader is the websocket upgrader.
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	// 简单允许所有跨域请求
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Client is a websocket client.
type Client struct {
	Username string
	Conn     *websocket.Conn
	Mutex    sync.Mutex
}

// clients is the map of all connected clients.
var (
	clients   = make(map[string]*Client)
	clientsMu sync.RWMutex
)

// HandleChat is the main function for the chat API
func HandleChat(w http.ResponseWriter, r *http.Request) {
	username := utils.GetCookieValue(r, "username")

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Upgrade error:", err)
		return
	}

	client := &Client{
		Username: username,
		Conn:     conn,
	}

	clientsMu.Lock()
	clients[username] = client
	clientsMu.Unlock()

	fmt.Printf("User %s connected.\n", username)

	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			fmt.Printf("Read error from %s: %v\n", username, err)
			break
		}
		fmt.Printf("Received from %s: %s\n", username, message)
	}

	clientsMu.Lock()
	delete(clients, username)
	clientsMu.Unlock()
	err = conn.Close()
	if err != nil {
		return
	}
	fmt.Printf("User %s disconnected.\n", username)
}
