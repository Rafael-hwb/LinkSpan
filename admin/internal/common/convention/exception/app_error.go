package exception

import "linkspan/admin/internal/common/convention/errorcode"

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
