package main

import (
	// "encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/hrrydgls/lego/controllers"
	// "github.com/hrrydgls/lego/models/responses"
	"github.com/hrrydgls/lego/services"
	// "github.com/hrrydgls/lego/services/auth"
)

func aboutHandler (w http.ResponseWriter, _ *http.Request) {
	hobbies := []string {"Football", "Hiking", "Musics"}
	fmt.Fprint(w, "<h1>About us!</h1>")

	fmt.Fprintln(w, "<p>My hobbies are: </p>")

	fmt.Fprintln(w, "<ul>")

	for _, hobby := range hobbies {
		fmt.Fprintf(w, "<li>%s</li>", hobby)
	}

	fmt.Fprintln(w, "</ul>")

}

func notFoundHandler (w http.ResponseWriter, r *http.Request) {
	services.ReturnResponse(w, services.NewNotFoundResponse())
}

func main() {
	r := chi.NewRouter()

	r.Get("/", controllers.HomeHandler)
	r.Get("/health", controllers.HealthHandler)
	r.Get("/about", aboutHandler)
	r.Get("/json", controllers.JsonController)
	r.Post("/register", controllers.RegisterController)
	r.Post("/login", controllers.LoginController)
	r.Get("/me", controllers.MeController)

	r.NotFound(notFoundHandler)


	slog.Info("Listenning to port 8000!")
	http.ListenAndServe(":8000", r)
}
