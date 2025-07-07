package geofence

import "math"

var geofenceLat = -6.2088
var geofenceLon = 106.8456
var radiusMeter = 50.0

func CheckInGeofence(lat, lon float64) bool {
	distance := Haversine(lat, lon, geofenceLat, geofenceLon)
	return distance <= radiusMeter
}

func Haversine(lat1, lon1, lat2, lon2 float64) float64 {
	const R = 6371000 // Earth radius in meters
	dLat := (lat2 - lat1) * math.Pi / 180
	dLon := (lon2 - lon1) * math.Pi / 180

	a := math.Sin(dLat/2)*math.Sin(dLat/2) +
		math.Cos(lat1*math.Pi/180)*math.Cos(lat2*math.Pi/180)*
			math.Sin(dLon/2)*math.Sin(dLon/2)

	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	return R * c
}
