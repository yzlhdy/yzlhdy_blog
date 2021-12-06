package helper

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
	Data    interface{} `json:"data"`
}

type EmptyObj struct{}

func BuildResponse(status int, message string, data interface{}) Response {
	return Response{
		Status:  status,
		Message: message,
		Data:    data,
		Error:   nil,
	}
}

func BuildErrorResponse(status int, message string, data interface{}, error interface{}) Response {
	return Response{
		Status:  status,
		Message: message,
		Data:    data,
		Error:   error,
	}
}
