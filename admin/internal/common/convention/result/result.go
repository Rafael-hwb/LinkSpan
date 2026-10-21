package result


const SuccessCode = "0"

//所有业务接口共用的相应结构

type Result[T any] struct {
	Code      string  `json:"code"`
	Message   string  `json:"message"`
	Data      T       `json:"data"`
	RequestID *string `json:"requestid"`
	Success   bool    `json:"success"`
}

func Success[T any](data T) Result[T] {
	return Result[T]{
		Code:    SuccessCode,
		Data:    data,
		Success: true,
	}
}

func Failure(code, message string) Result[any] {
	return Result[any]{
		Code:     code,
		Message:  message,
		Success:  code == SuccessCode,
	}
}