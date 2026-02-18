package controllers

import (
	"database/sql"
	"encoding/json"
	// "fmt"

	// "io"
	"log/slog"
	"net/http"

	"github.com/hrrydgls/lego/database"
	// "github.com/hrrydgls/lego/models"
	"github.com/hrrydgls/lego/models/requests"
	"github.com/hrrydgls/lego/models/responses"
)

func RegisterController(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		var data requests.RegisterRequest

		// body, _ := io.ReadAll(r.Body)

		json.NewDecoder(r.Body).Decode(&data)

		if data.Name == "" || data.Email == "" || data.Password == "" {
			w.WriteHeader(http.StatusBadRequest)
			response := map[string]string{
				"message": "Validation error!",
			}

			json.NewEncoder(w).Encode(response)
			return
		}
		db := database.DB()

		defer db.Close()

		var id uint
		var name string
		var email string

		err := db.QueryRow(
			"select id, name, email from users where email = $1",
			data.Email,
			).Scan(&id, &name, &email)
		var message string
		if err != nil {
			if err == sql.ErrNoRows {

				_, err := db.Exec(
					"insert into users (name, email, password) values ($1, $2, $3)",
					data.Name,
					data.Email,
					data.Password,
				)

				if err != nil {
					slog.Error("Error while creating the new user", "err", err)
					message = "Error while creating the new user!"
				} else {
					message = "We created a new user with this email!"
				}

			} else {
				slog.Error("Error in checking the user", "err", err)
				message = "Unkown error! check the logs..."
			}
		} else {
			message = "This email is already exist!"
		}

		json.NewEncoder(w).Encode(map[string]string{
			"message": message,
		})
		return

	} else {
		route := r.URL.Path
		method := r.Method
		response := responses.NotFound{
			Message: "Method is not supported!",
			Route:   route,
			Method:  method,
		}

		w.Header().Set("Accept", "application/json")
		w.WriteHeader(http.StatusMethodNotAllowed)

		json.NewEncoder(w).Encode(response)
		return
	}

}
