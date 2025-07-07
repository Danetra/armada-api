package model

import (
	"database/sql"
)

type LocationPayload struct {
	VehicleID string  `json:"vehicle_id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Timestamp int64   `json:"timestamp"`
}

// GetLatestLocation mengambil data lokasi terakhir
func GetLatestLocation(db *sql.DB, vehicleID string) (*LocationPayload, error) {
	query := `SELECT vehicle_id, latitude, longitude, timestamp FROM vehicle_locations WHERE vehicle_id = $1 ORDER BY timestamp DESC LIMIT 1`
	row := db.QueryRow(query, vehicleID)

	var loc LocationPayload
	err := row.Scan(&loc.VehicleID, &loc.Latitude, &loc.Longitude, &loc.Timestamp)
	return &loc, err
}

// GetLocationHistory mengambil data riwayat lokasi
func GetLocationHistory(db *sql.DB, vehicleID string, start, end int64) ([]LocationPayload, error) {
	query := `SELECT vehicle_id, latitude, longitude, timestamp FROM vehicle_locations WHERE vehicle_id = $1 AND timestamp BETWEEN $2 AND $3 ORDER BY timestamp ASC`
	rows, err := db.Query(query, vehicleID, start, end)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var locations []LocationPayload
	for rows.Next() {
		var loc LocationPayload
		if err := rows.Scan(&loc.VehicleID, &loc.Latitude, &loc.Longitude, &loc.Timestamp); err != nil {
			return nil, err
		}
		locations = append(locations, loc)
	}

	return locations, nil
}