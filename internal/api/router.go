package api

import (
	"armada-api/internal/rabbitmq"
	"database/sql"

	"github.com/gin-gonic/gin"
)

func StartServer(db *sql.DB, rmq *rabbitmq.RabbitMQ) {
	r := gin.Default()

	h := NewHandler(db)

	r.GET("/vehicles/:vehicle_id/location", h.GetLatestLocation)
	r.GET("/vehicles/:vehicle_id/history", h.GetLocationHistory)

	r.Run("0.0.0.0:8080")
}
