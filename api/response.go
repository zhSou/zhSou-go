package api

const (
	Success      = 0
	UnknownError = 1
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func BuildSuccessResponse(data interface{}) *Response {
	return &Response{
		Code: Success,
		Msg:  "Success",
		Data: data,
	}
}

func BuildErrorResponse(code int, msg string) *Response {
	return &Response{
		Code: code,
		Msg:  msg,
	}
}

func BuildUnknownErrorResponse(err error) *Response {
	return &Response{
		Code: UnknownError,
		Msg:  err.Error(),
	}
}
