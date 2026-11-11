package exception

import "linkspan/internal/common/convention/errorcode"

type ServiceError struct {
	Code    string
	Message string
	Cause   error
}

func NewServiceError(code errorcode.ErrorCode) *ServiceError {
	return &ServiceError{
		Code:    code.Code,
		Message: code.Message,
	}
}

func (e *ServiceError) Error() string {
	return e.Code + ": " + e.Message
}

func (e *ServiceError) Unwrap() error {
	return e.Cause
}

// WrapServiceError 保留原始原因，对外仍使用约定的错误码和消息。
func WrapServiceError(code errorcode.ErrorCode, cause error) *ServiceError {
	err := NewServiceError(code)
	err.Cause = cause
	return err
}
