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
	hub := notification.NewHub()
	go hub.Run()
	mux := http.NewServeMux()
	notification.SetupRoutes(mux, hub)
	go func() {
		for {
			time.Sleep(15 * time.Second)
			hub.BroadcastMessage("Periodic broadcast from server!")
		}
	}()

	log.Println("Server starting on", cfg.Port)
	if err := http.ListenAndServe(cfg.Port, mux); err != nil {
		log.Fatal("ListenAndServe error:", err)
	}
}
