package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"time"
)

type TemperatureResponse struct {
	Value       float64   `json:"value"`
	Unit        string    `json:"unit"`
	Timestamp   time.Time `json:"timestamp"`
	Location    string    `json:"location"`
	Status      string    `json:"status"`
	SensorID    string    `json:"sensor_id"`
	SensorType  string    `json:"sensor_type"`
	Description string    `json:"description"`
}

func main() {
	http.HandleFunc("/temperature", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		// Получаем параметры из query string
		location := r.URL.Query().Get("location")
		sensorID := r.URL.Query().Get("sensorId")

		// If no location is provided, use a default based on sensor ID
		if location == "" {
			switch sensorID {
			case "1":
				location = "Living Room"
			case "2":
				location = "Bedroom"
			case "3":
				location = "Kitchen"
			default:
				location = "Unknown"
			}
		}

		// If no sensor ID is provided, generate one based on location
		if sensorID == "" {
			switch location {
			case "Living Room":
				sensorID = "1"
			case "Bedroom":
				sensorID = "2"
			case "Kitchen":
				sensorID = "3"
			default:
				sensorID = "0"
			}
		}

		temperature := 18.0 + rand.Float64()*7.0

		response := TemperatureResponse{
			Value:       temperature,
			Unit:        "°C",
			Timestamp:   time.Now(),
			Location:    location,
			Status:      "active",
			SensorID:    sensorID,
			SensorType:  "temperature",
			Description: "Temperature sensor in " + location,
		}

		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, "Error encoding response", http.StatusInternalServerError)
			return
		}
	})

	http.ListenAndServe(":8081", nil)
}
