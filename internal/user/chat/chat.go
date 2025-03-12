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
	// Allow all origins.
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// Client is a websocket client.
type Client struct {
	Username  string
	Conn      *websocket.Conn
	Mutex     sync.Mutex
	ReplyChan chan []byte
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

	// Handle messages from the client.
	func(client *Client) {
		for {
			_, message, err := client.Conn.ReadMessage()
			if err != nil {
				client.ReplyChan <- nil
				break
			}
			// if client.ReplyChan is not nil, then send the message to the ReplyChan
			if client.ReplyChan != nil {
				select {
				case client.ReplyChan <- message:
				default:
					fmt.Printf("Received from %s (undelivered): %s\n", client.Username, message)
				}
			} else {
				fmt.Printf("Received from %s: %s\n", client.Username, message)
			}
		}
	}(client)

	clientsMu.Lock()
	delete(clients, username)
	clientsMu.Unlock()
	err = conn.Close()
	if err != nil {
		return
	}
	fmt.Printf("User %s disconnected.\n", username)
}
