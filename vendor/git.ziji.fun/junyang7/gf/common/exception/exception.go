package exception

type Exception struct {
	Code    int
	Message string
}

func New(code int, message string) *Exception {
	this := &Exception{
		Code:    code,
		Message: message,
	}
	return this
}
