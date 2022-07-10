package common

const (
	FailureCode    = -1        // 失败状态码
	FailureMessage = "failure" // 失败信息
	SuccessCode    = 0         // 成功状态码
	SuccessMessage = "success" // 成功消息
)

var (
	DefaultData = struct{}{} // 默认响应空对象
)
