package interceptor

import (
	"git.ziji.fun/junyang7/gf/common"
	"git.ziji.fun/junyang7/gf/common/exception"
)

type Interceptor struct {
	ok      bool
	code    int
	message string
	data    interface{}
}

func Insure(ok bool) *Interceptor {
	return &Interceptor{
		ok:      ok,
		code:    common.FailureCode,
		message: common.FailureMessage,
		data:    common.DefaultData,
	}
}
func (this *Interceptor) Exception(exception exception.Exception) *Interceptor {
	this.code = exception.Code
	this.message = exception.Message
	return this
}
func (this *Interceptor) Code(code int) *Interceptor {
	this.code = code
	return this
}
func (this *Interceptor) Message(message string) *Interceptor {
	this.message = message
	return this
}
func (this *Interceptor) Data(data interface{}) *Interceptor {
	this.data = data
	return this
}
func (this *Interceptor) Do() {
	if this.ok {
		return
	}
	response := common.NewResponse()
	response.Code = this.code
	response.Message = this.message
	response.Data = this.data
	panic(response)
}
