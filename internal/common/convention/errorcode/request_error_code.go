package errorcode

var (
	RequestBodyInvalid = ErrorCode{"A000002", "请求体格式错误"}
	ParameterInvalid   = ErrorCode{"A000003", "请求参数校验失败"}
	RequestBodyEmpty   = ErrorCode{"A000004", "请求体不能为空"}
	ServiceTimeout     = ErrorCode{"B000100", "系统执行超时"}
)
