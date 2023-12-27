package rest

type RestRespose struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Success bool        `json:"success"`
}

func Success(data interface{}) RestRespose {
	return RestRespose{
		Data:    data,
		Code:    200,
		Message: "success",
		Success: true,
	}
}

func Error(message string, code int) RestRespose {
	return RestRespose{
		Code:    code,
		Message: message,
		Data:    nil,
		Success: false,
	}
}
