package controllers

import (
	"net/http"

	"github.com/hrrydgls/lego/services"
)

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	services.ReturnResponse(w, services.NewSuccessResponse("App is up!", nil))
}