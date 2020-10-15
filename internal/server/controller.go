package server

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"time"

	"github.com/fri-go/internal/redis"
	"github.com/fri-go/types/sensor"
	"github.com/gorilla/mux"
)

// GetSensorDataBetweenDates : return sensor values between two dates for an airport
func GetSensorDataBetweenDates(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	const layoutISO = "2006-01-02"
	reqVars := mux.Vars(r)
	idAirport := reqVars["idAirport"]
	sensorType := reqVars["sensorType"]
	date1, err := time.Parse(layoutISO, reqVars["date1"])
	date2, err := time.Parse(layoutISO, reqVars["date2"])

	if err != nil {
		fmt.Println(err)
	}

	m := make(map[string][]float64)

	days := date2.Sub(date1).Hours() / 24

	if days == 0 {
		dateStr := time.Time.Format(date1, layoutISO)
		m[dateStr] = redis.GetSensorDataByDate(idAirport, dateStr, sensorType)
	} else {
		for i := 0; i < int(math.Abs(days))+1; i++ {
			dateStr := time.Time.Format(date1, layoutISO)
			m[dateStr] = redis.GetSensorDataByDate(idAirport, dateStr, sensorType)
			date1 = date1.AddDate(0, 0, 1)
		}
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(m)
}

// GetAverageValues : return the average of all 3 types of values for an airport and for a day
func GetAverageValues(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")

	reqVars := mux.Vars(r)
	idAirport := reqVars["idAirport"]
	date := reqVars["date"]

	averageTemperature := redis.GetAverageSensorValue(idAirport, date, "temperature")
	averagePressure := redis.GetAverageSensorValue(idAirport, date, "pressure")
	averageWind := redis.GetAverageSensorValue(idAirport, date, "wind")

	averageValues := sensor.Average{Temperature: averageTemperature, Pressure: averagePressure, Wind: averageWind}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(averageValues)
}
