package http

type response struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
	Time    int64       `json:"time"`
	Consume int64       `json:"consume"`
}

func newResponse() *response {
	return &response{
		Code:    0,
		Message: "success",
		Data:    struct{}{},
	}
}
