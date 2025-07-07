package api

import (
	"armada-api/internal/model"
	"database/sql"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Handler adalah struct yang menyimpan koneksi database
type Handler struct {
	DB *sql.DB
}

// NewHandler constructor (ini yang error tadi)
func NewHandler(db *sql.DB) *Handler {
	return &Handler{DB: db}
}

// GET /vehicles/:vehicle_id/location
func (h *Handler) GetLatestLocation(c *gin.Context) {
	vehicleID := c.Param("vehicle_id")

	location, err := model.GetLatestLocation(h.DB, vehicleID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Location not found"})
		return
	}

	c.JSON(http.StatusOK, location)
}

// GET /vehicles/:vehicle_id/history
func (h *Handler) GetLocationHistory(c *gin.Context) {
	vehicleID := c.Param("vehicle_id")
	startStr := c.Query("start")
	endStr := c.Query("end")

	start, err := strconv.ParseInt(startStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid start timestamp"})
		return
	}

	end, err := strconv.ParseInt(endStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid end timestamp"})
		return
	}

	history, err := model.GetLocationHistory(h.DB, vehicleID, start, end)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch history"})
		return
	}

	c.JSON(http.StatusOK, history)
}
