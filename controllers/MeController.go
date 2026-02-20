package controllers

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	// "github.com/hrrydgls/lego/models/responses"
	"github.com/hrrydgls/lego/services/auth"
)

func MeController(w http.ResponseWriter, r *http.Request) {
	authHeader := r.Header.Get("Authorization")

	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer "){
		http.Error(w, "Missing or invalid token format!", http.StatusUnauthorized)
		return
	}

	tokenString := strings.Split(authHeader, " ")[1]

	claims := auth.Claims{}


	token, err := jwt.ParseWithClaims(tokenString, &claims, func (token *jwt.Token)(interface{}, error){
		return auth.JwtSecret, nil
	}) 

	if err != nil || !token.Valid {
		http.Error(w, "Invalid token!", http.StatusUnauthorized)
		return
	}

	json.NewEncoder(w).Encode(map[string]uint {
		"user_id": claims.UserID,
	})
}