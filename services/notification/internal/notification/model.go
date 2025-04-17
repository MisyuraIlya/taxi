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
