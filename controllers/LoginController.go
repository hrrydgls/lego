package controllers

import (
	"database/sql"
	"encoding/json"

	// "fmt"
	"net/http"

	// "github.com/hrrydgls/lego/database"
	"github.com/hrrydgls/lego/database"
	"github.com/hrrydgls/lego/models/requests"
	"github.com/hrrydgls/lego/models/responses"
	// "github.com/hrrydgls/lego/models/responses/errors"
	"github.com/hrrydgls/lego/services"
	"github.com/hrrydgls/lego/services/auth"
)

// type JsonResponse struct {
// 	Message string `json:"message"`
// }

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

		services.ReturnResponse(w, services.NewMethodNotSupportedResponse())

		return
	}

	// check if email and pass are there
	var data requests.LoginRequest
	json.NewDecoder(r.Body).Decode(&data)

	if data.Email == "" || data.Password == "" {
		services.ReturnResponse(w, services.NewInvalidInputResponse("Both email and password are required!", nil))
		return
	}

	// check if the user(email) exist in db or not

	db := database.DB()

	query := "select id, name, email, password from users where email = $1"

	row := db.QueryRow(query, data.Email)
	var id uint 
	var name string
	var email string
	var password string

	err := row.Scan(&id, &name, &email, &password)

	if (err != nil) {
		if err == sql.ErrNoRows {
			services.ReturnResponse(w, services.NewInvalidInputResponse("User not found!", nil))
			return
		} else {
			services.ReturnResponse(w, services.NewServerErrorResponse())
			return
		}
	}


	if password == data.Password {
		token, _ := auth.GenerateToken(id)

		loginResponse := responses.LoginResponse {
			Id: id,
			Name: name,
			Email: email,
			Token: token,
		}

		// json.NewEncoder(w).Encode(loginResponse)

		services.ReturnResponse(w, services.NewSuccessResponse("Logged in successfuly!", loginResponse))

	} else {
		services.ReturnResponse(w, services.NewInvalidInputResponse("Invalid password!", nil))
	}



}