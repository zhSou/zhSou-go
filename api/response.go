package api

const (
	Success = iota
	UnknownError
	RequestFormatError
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
