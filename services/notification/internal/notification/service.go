package notification

import (
	"fmt"
	"log"
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

func NewHub() *Hub {
	return &Hub{
		Clients:    make(map[string]*Client),
		Register:   make(chan *Client),
		Unregister: make(chan *Client),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case client := <-h.Register:
			h.Clients[client.DriverID] = client
			log.Printf("Driver %s registered: %v", client.DriverID, client.Conn.RemoteAddr())

			if client.OnConnect != nil {
				go client.OnConnect(client)
			}

		case client := <-h.Unregister:
			if _, ok := h.Clients[client.DriverID]; ok {
				delete(h.Clients, client.DriverID)
				close(client.Send)
				log.Printf("Driver %s unregistered: %v", client.DriverID, client.Conn.RemoteAddr())

				if client.OnDisconnect != nil {
					go client.OnDisconnect(client)
				}
			}
		}
	}
}

func (h *Hub) BroadcastMessage(msg string) {
	for _, client := range h.Clients {
		select {
		case client.Send <- []byte(msg):
		default:
			close(client.Send)
			delete(h.Clients, client.DriverID)
		}
	}
}

func (h *Hub) SendAndWaitForResponse(driverID, message string) (string, error) {
	client, ok := h.Clients[driverID]
	if !ok {
		return "", fmt.Errorf("driver %s not connected", driverID)
	}

	client.PendingResponse = make(chan string, 1)

	select {
	case client.Send <- []byte(message):
	default:
		return "", fmt.Errorf("failed to send message to driver %s", driverID)
	}

	select {
	case resp := <-client.PendingResponse:
		return resp, nil
	case <-time.After(5 * time.Second):
		return "", fmt.Errorf("timeout waiting for driver %s response", driverID)
	}
}

func (c *Client) readPump() {
	defer func() {
		c.Hub.Unregister <- c
		c.Conn.Close()
	}()

	c.Conn.SetReadLimit(maxMessageSize)
	_ = c.Conn.SetReadDeadline(time.Now().Add(pongWait))
	c.Conn.SetPongHandler(func(string) error {
		_ = c.Conn.SetReadDeadline(time.Now().Add(pongWait))
		return nil
	})

	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				log.Printf("Unexpected close error for driver %s: %v", c.DriverID, err)
			}
			break
		}

		response := string(message)
		if c.PendingResponse != nil {
			select {
			case c.PendingResponse <- response:
			default:
			}
		}
	}
}

func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.Conn.Close()
	}()

	for {
		select {
		case msg, ok := <-c.Send:
			_ = c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				_ = c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			writer, err := c.Conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			_, _ = writer.Write(msg)

			n := len(c.Send)
			for i := 0; i < n; i++ {
				_, _ = writer.Write(newline)
				_, _ = writer.Write(<-c.Send)
			}

			if err := writer.Close(); err != nil {
				return
			}
		case <-ticker.C:
			_ = c.Conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.Conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}
