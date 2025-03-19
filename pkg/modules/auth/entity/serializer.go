package entity

type Response struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Meta    MetaData    `json:"meta"`
}

type MetaData struct {
	HttpStatus int `json:"http_status"`
	Total      int `json:"total"`
	Limit      int `json:"limit"`
	Offset     int `json:"offset"`
}

func LoginResponseSerializer(user any, token string, statusCode int, message string) Response {
	data := map[string]interface{}{
		"token": token,
		"user":  user,
	}

	return Response{
		Message: message,
		Data:    data,
		Meta: MetaData{
			HttpStatus: statusCode,
		},
	}
}

func UserSerializer(user *User, statusCode int, message string) Response {
	data := map[string]interface{}{
		"user": user,
	}

	return Response{
		Message: message,
		Data:    data,
		Meta: MetaData{
			HttpStatus: statusCode,
		},
	}
}
