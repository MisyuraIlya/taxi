package main

import (
	"geo-service/configs"
	"geo-service/internal/geo"
	"geo-service/pkg/redis"
	"log"
	"net/http"
)

func main() {
	config := configs.LoadConfig()
	redis := redis.NewRedisPool(config)
	repository := geo.NewRepository(redis)
	service := geo.NewService(repository)
	router := http.NewServeMux()
	geo.NewHandler(router, service)
	log.Println("Message service running on", config.APP_PORT)
	log.Fatal(http.ListenAndServe(config.APP_PORT, router))

}
