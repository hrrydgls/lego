package services

import (
	"encoding/json"
	"net/http"
)

type Response struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
	Success    bool   `json:"success"`
	Data interface{} `json:"data"`
}

func NewResponse(statusCode int, message string, success bool, data interface{}) Response {
	return Response{
		StatusCode: statusCode,
		Message:    message,
		Success:    success,
		Data: data,
	}
}

func NewSuccessResponse(message string, data interface{}) Response {
	return NewResponse(http.StatusAccepted, message, true, data)
}

func NewInvalidInputResponse(message string, data interface{}) Response {
	return NewResponse(http.StatusUnprocessableEntity, message, false, data)
}

func NewServerErrorResponse() Response {
	return NewResponse(http.StatusInternalServerError, "Server error!", false, nil)
}

func NewMethodNotSupportedResponse() Response {
	return NewResponse(http.StatusMethodNotAllowed, "This method is not supported!", false, nil)
}

func NewNotFoundResponse() Response {
	return NewResponse(http.StatusNotFound, "We couldn't find the route or resource!", false, nil)
}

func ReturnResponse(w http.ResponseWriter, response Response) {
	w.Header().Set("Accept", "applocation/json")
	w.WriteHeader(response.StatusCode)
	json.NewEncoder(w).Encode(response)
}
