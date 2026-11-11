package errorcode

import "linkspan/internal/common/convention/errorcode"

type ErrorCode = errorcode.ErrorCode

var (
	ClientError  = errorcode.ClientError
	ServiceError = errorcode.ServiceError
	RemoteError  = errorcode.RemoteError
)
