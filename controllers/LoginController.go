package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type JsonResponse struct {
	Message string `json:"message"`
}

func LoginController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Hi from login!")


	
	w.Header().Set("Accept", "application/json")


	w.WriteHeader(http.StatusAccepted)
	
	// fmt.Fprintf(w, "Hi from login!")



	// json.NewEncoder(w).Encode(map[string]string{
	// 	"message": "Hi from login!",
	// })

	response := JsonResponse{
		Message: "Hi from login again!",
	}

	json.NewEncoder(w).Encode(response)

}