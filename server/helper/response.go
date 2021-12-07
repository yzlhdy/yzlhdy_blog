package helper

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
	Data    interface{} `json:"data"`
}

type ResponseAll struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Error   interface{} `json:"error"`
	Data    interface{} `json:"data"`
	Total   int64       `json:"total"`
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

func BuildResponseAll(status int, message string, data interface{}, total int64, error interface{}) ResponseAll {
	return ResponseAll{
		Status:  status,
		Message: message,
		Data:    data,
		Error:   error,
		Total:   total,
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
