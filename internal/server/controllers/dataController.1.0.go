package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/fri-go/internal/server/models"
	"github.com/gorilla/mux"
)

func FilterDataByDate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	vars := mux.Vars(r)
	dateA, err := vars["dateA"]
	dateB, err := vars["dateB"]

	if err {
		log.Fatal(err)
	}

	models.FilterDataByDate(dateA, dateB)

	json.NewEncoder(w).Encode("end")
}

func GetAverageValues(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode("working")
}
