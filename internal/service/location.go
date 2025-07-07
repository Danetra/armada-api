package service

import (
	"armada-api/databases"
	"armada-api/internal/geofence"
	"armada-api/internal/model"
	"armada-api/internal/rabbitmq"
	"database/sql"
)

func HandleLocation(dbConn *sql.DB, rmq *rabbitmq.RabbitMQ, loc model.LocationPayload) {
	databases.InsertLocation(dbConn, loc)

	if geofence.CheckInGeofence(loc.Latitude, loc.Longitude) {
		rabbitmq.PublishGeofenceEvent(rmq, loc)
	}
}

