package controllers

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	Message string `json: "message"`
	Code int `json: "code"`
}

func JsonController(w http.ResponseWriter, r *http.Request) {
	response := Response{
		Message: "Hi from json endpoint",
		Code: 200,
	}

	json.NewEncoder(w).Encode(response)
}