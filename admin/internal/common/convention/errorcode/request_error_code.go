package errorcode

import "linkspan/internal/common/convention/errorcode"

// Go 适配层新增的请求及认证错误码；原有用户业务码保持不变。
var (
	RequestBodyInvalid = errorcode.RequestBodyInvalid
	ParameterInvalid   = errorcode.ParameterInvalid
	RequestBodyEmpty   = errorcode.RequestBodyEmpty
	NoUpdateFields     = ErrorCode{"A000005", "请至少提供一个需要修改的资料字段"}
	LoginRequired      = ErrorCode{"A000400", "请先登录"}
	SessionInvalid     = ErrorCode{"A000401", "登录状态无效或已过期"}
	CredentialsInvalid = ErrorCode{"A000402", "用户名或密码错误"}
	AccessDenied       = ErrorCode{"A000403", "无权修改其他用户的资料"}

	// 编号沿用原版公共错误码。
	UsernameInvalid = ErrorCode{"A000110", "用户名须为3至32位小写字母、数字或下划线"}
	PasswordInvalid = ErrorCode{"A000120", "密码不能超过72字节"}
	FlowLimitError  = ErrorCode{"A000300", "当前系统繁忙，请稍后再试"}
	ServiceTimeout  = errorcode.ServiceTimeout
)
