package helpers

//Json response representation
type Response struct {
	Message interface{} `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func BuildErrorResponse(message []string, data interface{}) Response {
	return Response{
		Message: message,
		Data:    data,
	}
}

func BuildResponse(message string, data interface{}) Response {
	return Response{
		Message: message,
		Data:    data,
	}
}
