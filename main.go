package main

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/hrrydgls/lego/controllers"
	"github.com/hrrydgls/lego/models/responses"
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
	notFoundResponse := responses.NotFound {
		Message: "Route not found!",
		Route: r.URL.Path,
		Method: r.Method,
	}
	json.NewEncoder(w).Encode(notFoundResponse)
}

func mainHandler(w http.ResponseWriter, r *http.Request) {

	path := r.URL.Path

	slog.Info("New request:", "path", path)

	switch path {
	case "/":
		controllers.HomeHandler(w, r)
	case "/health":
		controllers.HealthHandler(w, r)
	case "/about":
		aboutHandler(w, r)
	case "/json":
		controllers.JsonController(w, r)
	case "/register":
		controllers.RegisterController(w, r)
	case "/login":
		controllers.LoginController(w, r)
	case "/me":
		controllers.MeController(w, r)

		// auth.Test()
	default:
		notFoundHandler(w, r)
	}


}

func main() {
	slog.Info("App started...")
	http.HandleFunc("/", mainHandler)
	slog.Info("Listenning to port 8000!")
	http.ListenAndServe(":8000", nil)
}
