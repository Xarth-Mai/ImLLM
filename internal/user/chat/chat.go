package chat

import (
	"fmt"
	"github.com/Xarth-Mai/ImLLM/internal/utils"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
	"time"
)

const (
	pongWait   = 60 * time.Second
	pingPeriod = (pongWait * 9) / 10
	writeWait  = 10 * time.Second
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

	// set read deadline
	err = conn.SetReadDeadline(time.Now().Add(pongWait))
	if err != nil {
		return
	}
	// set write deadline
	conn.SetPongHandler(func(string) error {
		err := conn.SetReadDeadline(time.Now().Add(pongWait))
		if err != nil {
			return err
		}
		return nil
	})

	client := &Client{
		Username: username,
		Conn:     conn,
	}

	clientsMu.Lock()
	clients[username] = client
	clientsMu.Unlock()

	log.Printf("User %s connected.\n", username)

	// Handle Pings from the client.
	go func(c *Client) {
		ticker := time.NewTicker(pingPeriod)
		defer ticker.Stop()
		for {
			select {
			case <-ticker.C:
				// set write deadline
				err := c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
				if err != nil {
					log.Printf("Set write deadline error: %v\n", err)
				}
				// send ping message
				if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
					return
				}
			}
		}
	}(client)

	// Handle messages from the client.
	func(client *Client) {
		for {
			_, message, err := client.Conn.ReadMessage()
			if err != nil {
				if client.ReplyChan != nil {
					client.ReplyChan <- nil
				}
				break
			}
			// if client.ReplyChan is not nil, then send the message to the ReplyChan
			if client.ReplyChan != nil {
				select {
				case client.ReplyChan <- message:
				default:
					log.Printf("Received from %s (undelivered): %s\n", client.Username, message)
				}
			} else {
				log.Printf("Received from %s: %s\n", client.Username, message)
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
	log.Printf("User %s disconnected.\n", username)
}
