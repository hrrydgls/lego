package responses

type SuccessResponse struct {
	Message string `json:"message"`
	StatusCode int `json:"status_code"`
}


func NewSuccessResponse() SuccessResponse {
	return SuccessResponse {
		Message: "Your request was successful!",
		StatusCode: 200,
	}
}