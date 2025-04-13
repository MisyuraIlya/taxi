package main

import (
	"log"
	"net/http"
	"time"

	"notification-service/configs"
	"notification-service/internal/notification"
)

func main() {
	cfg := configs.NewConfig()

	// Create the notification hub.
	hub := notification.NewHub()
	// Start the hub's event loop in a separate goroutine.
	go hub.Run()

	// Create an HTTP multiplexer and register the endpoints.
	mux := http.NewServeMux()
	notification.SetupRoutes(mux, hub)

	// Example: a periodic broadcast (optional)
	go func() {
		for {
			time.Sleep(15 * time.Second)
			// This is just an example broadcast.
			// In your business logic, you will send targeted notifications.
			hub.BroadcastMessage("Periodic broadcast from server!")
		}
	}()

	log.Println("Server starting on", cfg.Port)
	if err := http.ListenAndServe(cfg.Port, mux); err != nil {
		log.Fatal("ListenAndServe error:", err)
	}
}
