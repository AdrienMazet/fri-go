package server

import (
	"encoding/json"
	"net/http"
    "time"
	"github.com/fri-go/internal/redis"
	"github.com/fri-go/types/sensor"
	"github.com/gorilla/mux"
)

const layoutISO = "2006-01-02";

// conf := configuration.LoadConfiguration()

// for i := 0; i < len(conf.Airports); i++ {
// 	airport := conf.Airports[i]
// 	fmt.Println(redis.GetAirportData(airport))
// 	fmt.Println(".")
// }

// GetSensorDataBetweenDates : return sensor values between two dates
func GetSensorDataBetweenDates(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	reqVars := mux.Vars(r)
	idAirport := reqVars["idAirport"]
	sensorType := reqVars["sensorType"]
	date1 := reqVars["date1"]
	date2 := reqVars["date2"]

    var m map[string]float64

    dateA := time.Parse(layoutISO, date1)
    dateB := time.Parse(layoutISO, date2)

    for (dateA.Before(dateB)) {
        dateStr := string(dateA.String()[0:10]);
    	key, value := redis.GetSensorDataByDate(idAirport, dateStr, sensorType)
    	m[key] = value
        dateA.AddDate(0, 0, 1)
    }

	json.NewEncoder(w).Encode(averageValues)

	// vars := mux.Vars(r)
	// dateA, err := vars["dateA"]
	// dateB, err := vars["dateB"]

	// if err {
	// 	log.Fatal(err)
	// }

	// client := redis.NewClient(&redis.Options{
	// 	Addr:     "localhost:6379",
	// 	Password: "",
	// 	DB:       0,
	// })

	// ctx := client.Context()

	// startDate, _ := time.Parse(layoutISO, dateA)
	// endDate, _ := time.Parse(layoutISO, dateB)
	// currentDate := startDate
	// for currentDate.Before(endDate) {
	// 	fmt.Println(startDate)
	// 	val, err := client.Get(ctx, dateA).Result()
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// 	fmt.Println(val)
	// 	currentDate = currentDate.Add(24 * time.Hour)
	// }

	json.NewEncoder(w).Encode("end")
}

// GetAverageValues : return the average of all 3 types of values for an airport and for a day
func GetAverageValues(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	reqVars := mux.Vars(r)
	idAirport := reqVars["idAirport"]
	date := reqVars["date"]

	averageTemperature := redis.GetAverageSensorValue(idAirport, date, "temperature")
	averagePressure := redis.GetAverageSensorValue(idAirport, date, "pressure")
	averageWind := redis.GetAverageSensorValue(idAirport, date, "wind")

	averageValues := sensor.Average{Temperature: averageTemperature, Pressure: averagePressure, Wind: averageWind}

	json.NewEncoder(w).Encode(averageValues)
}
