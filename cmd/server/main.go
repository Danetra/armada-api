package server

import (
	"armada-api/databases"
	"armada-api/internal/api"
	"armada-api/internal/mqtt"
	"armada-api/internal/rabbitmq"
	"fmt"
)


func Run() {
	db := databases.InitDB()
	defer db.Close()

	rmq := rabbitmq.InitRabbitMQ()
	defer rmq.Close()

	fmt.Println("Successfully connected!")

	// Jalankan MQTT subscriber
	go mqtt.StartSubscriber(db, rmq)

	// Jalankan Geofence Worker
	go rabbitmq.StartGeofenceWorker(rmq)

	// Jalankan HTTP API
	api.StartServer(db, rmq)
}