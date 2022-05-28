package api

const (
	Success      = 0
	UnknownError = 1
)

type ApiResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func BuildSuccessResponse(data interface{}) *ApiResponse {
	return &ApiResponse{
		Code: Success,
		Msg:  "Success",
		Data: data,
	}
}

func BuildErrorResponse(code int, msg string) *ApiResponse {
	return &ApiResponse{
		Code: code,
		Msg:  msg,
	}
}

func BuildUnknownErrorResponse(err error) *ApiResponse {
	return &ApiResponse{
		Code: UnknownError,
		Msg:  err.Error(),
	}
}
