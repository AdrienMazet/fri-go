package server

import (
	"github.com/gorilla/mux"
)

// InitializeRouter : create a new router with 2 routes
func InitializeRouter() *mux.Router {
	// redirect /cars/ to /cars
	router := mux.NewRouter().StrictSlash(true)

	router.Methods("GET").Path("/{idAirport}/filter/{sensorType}/{date1}/{date2}").Name("getSensorDataBetweenDates").HandlerFunc(GetSensorDataBetweenDates)
	router.Methods("GET").Path("/{idAirport}/{date}/results/average").Name("getAverageValues").HandlerFunc(GetAverageValues)

	return router
}
