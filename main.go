package main

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/hrrydgls/lego/controllers"
)

func main() {
	r := chi.NewRouter()

	r.Get("/", controllers.HomeHandler)
	r.Get("/health", controllers.HealthHandler)
	r.Get("/about", controllers.AboutHandler)
	r.Get("/json", controllers.JsonController)
	r.Post("/register", controllers.RegisterController)
	r.Post("/login", controllers.LoginController)
	r.Get("/me", controllers.MeController)

	r.NotFound(controllers.NotFoundHandler)

	slog.Info("Listenning to port 8000!")
	http.ListenAndServe(":8000", r)
}
