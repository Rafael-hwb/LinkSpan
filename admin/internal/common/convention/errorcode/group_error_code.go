package errorcode

var (
	GroupNameInvalid = ErrorCode{
		Code:    "A000500",
		Message: "分组名称必须为1至64个字符",
	}

	GroupLimitExceeded = ErrorCode{
		Code:    "A000501",
		Message: "已达到分组数量上限",
	}

	GroupGIDGenerateFailed = ErrorCode{
		Code:    "B000500",
		Message: "分组标识生成失败，请稍后再试",
	}
)
