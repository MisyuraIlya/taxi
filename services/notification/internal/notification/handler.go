package notification

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func WsHandler(hub *Hub, w http.ResponseWriter, r *http.Request) {
	driverID := r.URL.Query().Get("driver_id")
	clientID := r.URL.Query().Get("client_id")

	if driverID == "" && clientID == "" {
		http.Error(w, "Missing driver_id or client_id", http.StatusBadRequest)
		return
	}

	id := driverID
	role := "driver"
	if clientID != "" {
		id = clientID
		role = "client"
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket upgrade error for %s %s: %v", role, id, err)
		return
	}

	client := &Client{
		DriverID: id,
		Hub:      hub,
		Conn:     conn,
		Send:     make(chan []byte, 256),
		OnConnect: func(c *Client) {
			log.Printf("[Handler] %s connected: %s", role, c.DriverID)
		},
		OnDisconnect: func(c *Client) {
			log.Printf("[Handler] %s disconnected: %s", role, c.DriverID)
		},
	}
	hub.Register <- client

	go client.readPump()
	go client.writePump()
}

func TargetedNotifyHandler(hub *Hub, w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	type notifyRequest struct {
		ClientID string `json:"client_id,omitempty"`
		DriverID string `json:"driver_id,omitempty"`
		Message  string `json:"message"`
	}

	var req notifyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request: "+err.Error(), http.StatusBadRequest)
		return
	}

	var id, role string
	if req.ClientID != "" {
		id = req.ClientID
		role = "client"
	} else if req.DriverID != "" {
		id = req.DriverID
		role = "driver"
	} else {
		http.Error(w, "Missing client_id or driver_id", http.StatusBadRequest)
		return
	}

	if req.Message == "" {
		http.Error(w, "Missing message", http.StatusBadRequest)
		return
	}

	response, err := hub.SendAndWaitForResponse(id, req.Message)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error sending to %s %s: %v", role, id, err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("%s %s responded: %s", role, id, response)))
}

func PushClientsHandler(hub *Hub, w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Expect exactly one of client_id or driver_id (but we'll only push to client here)
	type pushRequest struct {
		ClientID string `json:"client_id"`
		Message  string `json:"message"`
	}

	var req pushRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request: "+err.Error(), http.StatusBadRequest)
		return
	}
	if req.ClientID == "" || req.Message == "" {
		http.Error(w, "Missing client_id or message", http.StatusBadRequest)
		return
	}

	client, ok := hub.Clients[req.ClientID]
	if !ok {
		http.Error(w, "Client "+req.ClientID+" not connected", http.StatusNotFound)
		return
	}

	// Non‑blocking push
	select {
	case client.Send <- []byte(req.Message):
	default:
		// If their Send channel is full, drop the message or handle cleanup:
		close(client.Send)
		delete(hub.Clients, req.ClientID)
		http.Error(w, "Client send channel full; disconnected", http.StatusServiceUnavailable)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Message pushed to client " + req.ClientID))
}

func SetupRoutes(mux *http.ServeMux, hub *Hub) {
	mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		WsHandler(hub, w, r)
	})
	mux.HandleFunc("/notify/driver", func(w http.ResponseWriter, r *http.Request) {
		TargetedNotifyHandler(hub, w, r)
	})
	mux.HandleFunc("/notify/clients", func(w http.ResponseWriter, r *http.Request) {
		PushClientsHandler(hub, w, r)
	})
}
