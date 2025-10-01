package main

import (
	"encoding/json"
	"math/rand"
	"net/http"
)

type TemperatureResponse struct {
	Location    string  `json:"location"`
	SensorID    string  `json:"sensorId"`
	Temperature float64 `json:"temperature"`
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
		
		// Генерируем случайную температуру от 18.0 до 25.0
		temperature := 18.0 + rand.Float64()*7.0
		
		// Возвращаем JSON с данными
		response := TemperatureResponse{
			Location:    location,
			SensorID:    sensorID,
			Temperature: temperature,
		}
		
		json.NewEncoder(w).Encode(response)
	})
	
	http.ListenAndServe(":8081", nil)
}