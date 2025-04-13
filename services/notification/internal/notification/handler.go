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
	// For production, ensure proper origin checking.
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// WsHandler upgrades the HTTP connection to a WebSocket.
// It expects a "driver_id" parameter in the query string.
func WsHandler(hub *Hub, w http.ResponseWriter, r *http.Request) {
	driverID := r.URL.Query().Get("driver_id")
	if driverID == "" {
		http.Error(w, "Missing driver_id", http.StatusBadRequest)
		return
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("WebSocket upgrade error for driver %s: %v", driverID, err)
		return
	}

	client := &Client{
		DriverID: driverID,
		Hub:      hub,
		Conn:     conn,
		Send:     make(chan []byte, 256),
	}
	hub.Register <- client

	go client.readPump()
	go client.writePump()
}

// TargetedNotifyHandler handles POST requests to /notify/driver.
// It expects a JSON payload with "driver_id" and "message" fields.
// It sends the notification to the specified driver's WebSocket and waits up to 5 seconds for a response.
func TargetedNotifyHandler(hub *Hub, w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	type notifyRequest struct {
		DriverID string `json:"driver_id"`
		Message  string `json:"message"`
	}

	var req notifyRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request: "+err.Error(), http.StatusBadRequest)
		return
	}
	if req.DriverID == "" || req.Message == "" {
		http.Error(w, "Missing driver_id or message", http.StatusBadRequest)
		return
	}

	response, err := hub.SendAndWaitForResponse(req.DriverID, req.Message)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error: %v", err), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Driver %s responded: %s", req.DriverID, response)))
}

// SetupRoutes registers the HTTP endpoints for WebSocket connections and targeted notifications.
func SetupRoutes(mux *http.ServeMux, hub *Hub) {
	mux.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		WsHandler(hub, w, r)
	})
	mux.HandleFunc("/notify/driver", func(w http.ResponseWriter, r *http.Request) {
		TargetedNotifyHandler(hub, w, r)
	})
}
