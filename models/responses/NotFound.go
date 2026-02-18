package responses

type NotFound struct {
	Message string `json:"message"`
	Route string `json:"route"`
	Method string `json:"method"`
}