package notification

import "github.com/gorilla/websocket"

// Client represents a single WebSocket connection associated with a driver.
type Client struct {
	DriverID        string          // Unique identifier for the driver.
	Hub             *Hub            // Reference to the hub.
	Conn            *websocket.Conn // The WebSocket connection.
	Send            chan []byte     // Outbound messages queue.
	PendingResponse chan string     // Channel to receive driver's response.
}

// Hub maintains the set of active driver clients.
type Hub struct {
	Clients    map[string]*Client // Maps a driver's ID to its client.
	Register   chan *Client       // Channel for registering new clients.
	Unregister chan *Client       // Channel for unregistering clients.
}
