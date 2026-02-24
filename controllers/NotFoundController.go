package controllers

import (
	"net/http"

	"github.com/hrrydgls/lego/services"
)


func NotFoundHandler (w http.ResponseWriter, r *http.Request) {
	services.ReturnResponse(w, services.NewNotFoundResponse())
}