package errorcode

var (
	UserNull = ErrorCode{
		Code:    "B000200",
		Message: "用户记录不存在",
	}

	UserNameExist = ErrorCode{
		Code:    "B000201",
		Message: "用户名已存在",
	}

	UserExist = ErrorCode{
		Code:    "B000202",
		Message: "用户记录已存在",
	}

	UserSaveError = ErrorCode{
		Code:    "B000203",
		Message: "用户记录新增失败",
	}
)
