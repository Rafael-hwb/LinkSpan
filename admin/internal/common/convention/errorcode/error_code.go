package errorcode

//一项业务错误约定
type ErrorCode struct {
	Code string
	Message string
}

var(
	ClientError = ErrorCode {
		Code: "A000001",
		Message: "用户端错误",
	}

	ServiceError = ErrorCode {
		Code: "B000001",
		Message: "系统执行错误",
	}

	UserNull = ErrorCode {
		Code: "B000200",
		Message: "用户记录不存在",
	}
)