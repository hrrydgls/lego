package controllers

import (
	// "encoding/json"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	// "github.com/hrrydgls/lego/models/responses"
	"github.com/hrrydgls/lego/services"
	"github.com/hrrydgls/lego/services/auth"
)

func MeController(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")

	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		services.ReturnResponse(w, services.NewInvalidInputResponse("Missing token or invalid token format!", nil))
		return
	}

	tokenString := strings.Split(authHeader, " ")[1]

	claims := auth.Claims{}

	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return auth.JwtSecret, nil
	})

	if err != nil || !token.Valid {
		services.ReturnResponse(w, services.NewInvalidInputResponse("Invalid token!", nil))
		return
	}

	services.ReturnResponse(w, services.NewSuccessResponse("Logged in successfuly!", map[string]uint{
		"user_id": claims.UserID,
	}))
}
