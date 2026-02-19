package controllers

import (
	"encoding/json"
	// "fmt"
	"net/http"

	"github.com/hrrydgls/lego/database"
	"github.com/hrrydgls/lego/models/requests"
	"github.com/hrrydgls/lego/models/responses"
	"github.com/hrrydgls/lego/models/responses/errors"
)

type JsonResponse struct {
	Message string `json:"message"`
}

func LoginController(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("Hi from login!")
	
	// w.Header().Set("Accept", "application/json")


	// w.WriteHeader(http.StatusAccepted)
	
	// fmt.Fprintf(w, "Hi from login!")

	// json.NewEncoder(w).Encode(map[string]string{
	// 	"message": "Hi from login!",
	// })

	// response := JsonResponse{
	// 	Message: "Hi from login again!",
	// }

	// json.NewEncoder(w).Encode(response)


	// check if the request is post
	if r.Method != http.MethodPost {
		response := responses.NotFound{
			Message: "Method is not supported!",
			Route: r.URL.Path,
			Method: r.Method,
		}

		json.NewEncoder(w).Encode(response)
		return
	}

	// check if email and pass are there
	var data requests.LoginRequest
	json.NewDecoder(r.Body).Decode(&data)

	if data.Email == "" || data.Password == "" {
		json.NewEncoder(w).Encode(errors.NewValidationError())
		return
	}


	json.NewEncoder(w).Encode(responses.NewSuccessResponse())


}