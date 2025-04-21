package notification

import "github.com/gorilla/websocket"

type Client struct {
	DriverID        string
	Hub             *Hub
	Conn            *websocket.Conn
	Send            chan []byte
	PendingResponse chan string

	OnConnect    func(c *Client)
	OnDisconnect func(c *Client)
}
type Hub struct {
	Clients    map[string]*Client
	Register   chan *Client
	Unregister chan *Client
}

type Message struct {
	Type string `json:"type"`
	Data string `json:"data"`
}

type ClientMessage struct {
	ClientID string  `json:"client_id"`
	Message  Message `json:"message"`
}

type DriverMessage struct {
	ClientID string  `json:"client_id,omitempty"`
	DriverID string  `json:"driver_id,omitempty"`
	Message  Message `json:"message"`
}
