package rabbitmq

import (
	"log"
)

func StartGeofenceWorker(rmq *RabbitMQ) {
	// Pastikan queue & binding sudah ada
	queueName := "geofence_alerts"

	// Declare queue (jika belum ada)
	q, err := rmq.Ch.QueueDeclare(
		queueName, // name
		true,      // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
	)
	if err != nil {
		log.Fatal("Failed to declare queue:", err)
	}

	// Bind queue ke exchange
	err = rmq.Ch.QueueBind(
		q.Name,         // queue name
		"",             // routing key (kosong untuk fanout)
		"fleet.events", // exchange
		false,
		nil,
	)
	if err != nil {
		log.Fatal("Failed to bind queue:", err)
	}

	// Start consuming
	msgs, err := rmq.Ch.Consume(
		q.Name,
		"",
		true,  // auto-ack
		false, // not exclusive
		false,
		false,
		nil,
	)
	if err != nil {
		log.Fatal("Failed to register consumer:", err)
	}

	log.Println(" [*] Waiting for geofence events...")

	for msg := range msgs {
		log.Printf(" [x] Received geofence alert: %s", msg.Body)

		// Kamu bisa tambahkan logika simpan ke DB, send HTTP request, dsb. di sini
	}
}
