package main

import (
	"github.com/gorilla/mux"
)

func InitializeRouter() *mux.Router {
	// StrictSlash is true => redirect /cars/ to /cars
	router := mux.NewRouter().StrictSlash(true)

	router.Methods("GET").Path("/{dataType}/filter/date").Name("filterDataByDate").HandlerFunc(controllers.filterDataByDate)
	router.Methods("GET").Path("/results/average").Name("getAverageValues").HandlerFunc(controllers.getAverageValues)
	return router
}
