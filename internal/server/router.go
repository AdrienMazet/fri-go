package server

import (
	"github.com/fri-go/internal/server/controllers"
	"github.com/gorilla/mux"
)

// InitializeRouter : create a new router with 2 routes
func InitializeRouter() *mux.Router {
	// redirect /cars/ to /cars
	router := mux.NewRouter().StrictSlash(true)

	router.Methods("GET").Path("{IDAirport}/filter/{SensorType}/{date1}/{date2}").Name("getSensorDataBetweenDates").HandlerFunc(controllers.GetSensorDataBetweenDates)
	router.Methods("GET").Path("{IDAirport}/results/average").Name("getAverageValues").HandlerFunc(controllers.GetAverageValues)

	return router
}
