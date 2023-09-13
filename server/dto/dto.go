package dto

import "net/http"

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func GetErrorResponse(message string) Response {
	return Response{
		Code:    http.StatusInternalServerError,
		Message: message,
		Data:    nil,
	}
}

func GetSuccessResponse(data interface{}) Response {
	return Response{
		Code:    http.StatusOK,
		Message: "ok",
		Data:    data,
	}
}
