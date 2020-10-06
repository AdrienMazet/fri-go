package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/synbioz/go_api/models"
)

func filterDataByDate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(models.filterDataByDate(string dateA, string dateB))
}

func getAverageValues(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(models.getAverageValues())
}
