package utils

import "net/http"

type CommonResponse struct {
	Code   int    `json:"code"`
	Status string `json:"status"`
}

type ResponseWithMessage struct {
	CommonResponse
	Message string `json:"message"`
}

type ResponseWithData struct {
	CommonResponse
	Data interface{} `json:"data"`
}

func SuccessResponse() (int, CommonResponse) {
	return http.StatusOK, CommonResponse{
		Code:   http.StatusOK,
		Status: "success",
	}
}

func SuccessResponseWithData(data interface{}) (int, ResponseWithData) {
	code := http.StatusOK
	return code, ResponseWithData{
		CommonResponse: CommonResponse{
			Code:   code,
			Status: "success",
		},
		Data: data,
	}
}

func ClientErrorResponse(code int, message string) (int, ResponseWithMessage) {
	return code, ResponseWithMessage{
		CommonResponse: CommonResponse{
			Code:   code,
			Status: "fail",
		},
		Message: message,
	}
}

func ServerErrorResponse() (int, ResponseWithMessage) {
	return http.StatusInternalServerError, ResponseWithMessage{
		CommonResponse: CommonResponse{
			Code:   http.StatusInternalServerError,
			Status: "error",
		},
		Message: "sorry, there was a failure on our server",
	}
}
