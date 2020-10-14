package main

import (
	"github.com/fri-go/internal/server/controllers"
	"github.com/gorilla/mux"
)

func InitializeRouter() *mux.Router {
	// StrictSlash is true => redirect /cars/ to /cars
	router := mux.NewRouter().StrictSlash(true)

	router.Methods("GET").Path("/{dataType}/filter/date").Name("filterDataByDate").HandlerFunc(controllers.FilterDataByDate)
	router.Methods("GET").Path("/results/average").Name("getAverageValues").HandlerFunc(controllers.GetAverageValues)
	return router
}
