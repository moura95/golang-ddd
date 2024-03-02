package util

type Response struct {
	Error interface{} `json:"error"`
	Data  interface{} `json:"data"`
}

func SuccessResponse(data interface{}) Response {
	return Response{
		Data:  data,
		Error: "",
	}
}

func ErrorResponse(error string) Response {
	return Response{
		Data:  "",
		Error: error,
	}
}
