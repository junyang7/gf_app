package common

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Time    int64       `json:"time"`
	Consume int64       `json:"consume"`
}

func NewResponse() *Response {
	return &Response{
		Code:    SuccessCode,
		Message: SuccessMessage,
		Data:    DefaultData,
	}
}
