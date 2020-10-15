package server

import (
	"encoding/json"
<<<<<<< HEAD
	"fmt"
	"math"
	"net/http"
<<<<<<< HEAD
	"time"
=======
	"net/http"
>>>>>>> b89d3dd... Route for average done

=======
    "time"
>>>>>>> 2d3e3f0... implement getsensordatabetweendates
	"github.com/fri-go/internal/redis"
	"github.com/fri-go/types/sensor"
	"github.com/gorilla/mux"
)

<<<<<<< HEAD
<<<<<<< HEAD
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
=======
=======
const layoutISO = "2006-01-02";

>>>>>>> 2d3e3f0... implement getsensordatabetweendates
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
>>>>>>> b89d3dd... Route for average done
}

// GetAverageValues : return the average of all 3 types of values for an airport and for a day
func GetAverageValues(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
<<<<<<< HEAD
=======
	w.WriteHeader(http.StatusOK)
>>>>>>> b89d3dd... Route for average done

	reqVars := mux.Vars(r)
	idAirport := reqVars["idAirport"]
	date := reqVars["date"]

	averageTemperature := redis.GetAverageSensorValue(idAirport, date, "temperature")
	averagePressure := redis.GetAverageSensorValue(idAirport, date, "pressure")
	averageWind := redis.GetAverageSensorValue(idAirport, date, "wind")

	averageValues := sensor.Average{Temperature: averageTemperature, Pressure: averagePressure, Wind: averageWind}

<<<<<<< HEAD
	w.WriteHeader(http.StatusOK)
=======
>>>>>>> b89d3dd... Route for average done
	json.NewEncoder(w).Encode(averageValues)
}
