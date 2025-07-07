package rabbitmq

import (
	"armada-api/internal/model"
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

type RabbitMQ struct {
	Conn *amqp.Connection
	Ch   *amqp.Channel
}

func InitRabbitMQ() *RabbitMQ {
	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		log.Fatal("Failed to connect to RabbitMQ:", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal("Failed to open a channel:", err)
	}

	err = ch.ExchangeDeclare(
		"fleet.events",
		"fanout",       
		true,           
		false,          
		false,          
		false,          
		nil,            
	)
	if err != nil {
		log.Fatal("Failed to declare exchange:", err)
	}

	return &RabbitMQ{Conn: conn, Ch: ch}
}

func PublishGeofenceEvent(rmq *RabbitMQ, loc model.LocationPayload) {
	// Buat message payload
	msg := map[string]interface{}{
		"vehicle_id": loc.VehicleID,
		"event":      "geofence_entry",
		"location": map[string]float64{
			"latitude":  loc.Latitude,
			"longitude": loc.Longitude,
		},
		"timestamp": loc.Timestamp,
	}

	body, err := json.Marshal(msg)
	if err != nil {
		log.Println("Failed to marshal geofence event:", err)
		return
	}

	// Publish message ke exchange
	err = rmq.Ch.Publish(
		"fleet.events", // exchange
		"",             // routing key (kosong kalau pakai fanout)
		false,          // mandatory
		false,          // immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		log.Println("Failed to publish message:", err)
	}
}

func (rmq *RabbitMQ) Close() {
	if rmq.Ch != nil {
		_ = rmq.Ch.Close()
	}
	if rmq.Conn != nil {
		_ = rmq.Conn.Close()
	}
}
