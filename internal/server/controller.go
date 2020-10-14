package server

import (
	"encoding/json"
	"net/http"

	"github.com/fri-go/internal/redis"
	"github.com/fri-go/types/sensor"
	"github.com/gorilla/mux"
)

// conf := configuration.LoadConfiguration()

// for i := 0; i < len(conf.Airports); i++ {
// 	airport := conf.Airports[i]
// 	fmt.Println(redis.GetAirportData(airport))
// 	fmt.Println(".")
// }

// GetSensorDataBetweenDates : return sensor values between to dates
func GetSensorDataBetweenDates(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

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
