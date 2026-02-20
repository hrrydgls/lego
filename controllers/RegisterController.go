package controllers

import (
	"database/sql"
	"encoding/json"

	// "fmt"

	// "io"
	"log/slog"
	"net/http"

	"github.com/hrrydgls/lego/database"
	"github.com/hrrydgls/lego/services"
	// "github.com/hrrydgls/lego/models"
	"github.com/hrrydgls/lego/models/requests"
	// "github.com/hrrydgls/lego/models/responses"
)

func RegisterController(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodPost {

		var data requests.RegisterRequest

		// body, _ := io.ReadAll(r.Body)

		json.NewDecoder(r.Body).Decode(&data)

		if data.Name == "" || data.Email == "" || data.Password == "" {
			services.ReturnResponse(w, services.NewInvalidInputResponse("All fields are required!", nil))
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
					services.ReturnResponse(w, services.NewServerErrorResponse())
					return
				} else {
					services.ReturnResponse(w, services.NewSuccessResponse("We created a new user with this email!", data))
					return
				}

			} else {
				slog.Error("Error in checking the user", "err", err)
				services.ReturnResponse(w, services.NewServerErrorResponse())
				return
			}
		} else {
			message = "This email is already exist!"
			services.ReturnResponse(w, services.NewInvalidInputResponse(message, nil))
			return
		}

	} else {
		services.ReturnResponse(w, services.NewMethodNotSupportedResponse())

		return
	}

}
