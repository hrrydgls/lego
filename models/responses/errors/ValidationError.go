package errors


type ValidationError struct {
	Message string `json:"message"`
	StatusCode int `json:"status_code"`
}


func NewValidationError() ValidationError {
	return ValidationError{
		Message: "invalid input",
		StatusCode: 422,
	}
}