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

	GroupUnavailable = ErrorCode {
		Code:    "A000502",
		Message: "分组不可操作",
	}

	GroupGIDGenerateFailed = ErrorCode{
		Code:    "B000500",
		Message: "分组标识生成失败，请稍后再试",
	}
)
