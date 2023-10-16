package models

type response struct {
	Data       interface{} `json:"data"`
	StatusCode int         `json:"statusCode"`
	Error      string      `json:"error"`
	Message    string      `json:"message"`
}

func NewStandardResponse(data interface{}, statusCode int, err string, message string) response {
	return response{Data: data, StatusCode: statusCode, Error: err, Message: message}
}
