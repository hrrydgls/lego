package controllers

import (
	"encoding/json"
	"net/http"
)

type HealthResponse struct {
	Message string `json:"message"`
}

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	response := HealthResponse{
		Message: "The app is up!",
	}

	json.NewEncoder(w).Encode(response)
}