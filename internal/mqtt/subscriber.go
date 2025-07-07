package mqtt

import (
	"armada-api/internal/model"
	"armada-api/internal/rabbitmq"
	"armada-api/internal/service"
	"database/sql"
	"encoding/json"
	"fmt"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// Tambahkan parameter rmq *rabbitmq.RabbitMQ
func StartSubscriber(dbConn *sql.DB, rmq *rabbitmq.RabbitMQ) {
	opts := mqtt.NewClientOptions().AddBroker("tcp://mosquitto:1883")
	opts.SetClientID("fleet-subscriber")
	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	topic := "/fleet/vehicle/+/location"
	client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		var loc model.LocationPayload
		if err := json.Unmarshal(msg.Payload(), &loc); err != nil {
			fmt.Println("Invalid JSON:", err)
			return
		}

		// Pass db dan rabbitMQ ke HandleLocation
		service.HandleLocation(dbConn, rmq, loc)
	})

	select {} // keep subscriber running
}
