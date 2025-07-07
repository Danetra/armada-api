package mqtt

import (
	"armada-api/internal/model"
	"armada-api/internal/rabbitmq"
	"armada-api/internal/service"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

func StartSubscriber(dbConn *sql.DB, rmq *rabbitmq.RabbitMQ) {
	// Ambil broker URL dari ENV
	brokerURL := os.Getenv("MQTT_BROKER_URL")
	if brokerURL == "" {
		brokerURL = "tcp://localhost:1883"
	}

	opts := mqtt.NewClientOptions().
		AddBroker(brokerURL).
		SetClientID("fleet-subscriber").
		SetAutoReconnect(true).
		SetConnectRetry(true).
		SetConnectRetryInterval(5 * time.Second)

	// Optional: Log saat connect
	opts.OnConnect = func(c mqtt.Client) {
		fmt.Println("Connected to MQTT broker:", brokerURL)
	}

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(fmt.Sprintf("MQTT connection error: %v", token.Error()))
	}

	// Subscribe to topic
	topic := "/fleet/vehicle/+/location"
	if token := client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		var loc model.LocationPayload
		if err := json.Unmarshal(msg.Payload(), &loc); err != nil {
			fmt.Println("Invalid JSON payload:", err)
			return
		}

		// Process location data
		service.HandleLocation(dbConn, rmq, loc)
	}); token.Wait() && token.Error() != nil {
		panic(fmt.Sprintf("MQTT subscribe error: %v", token.Error()))
	}

	fmt.Println("Subscribed to topic:", topic)

	select {} // keep subscriber running forever
}
