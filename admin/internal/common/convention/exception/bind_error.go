package exception

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"

	"linkspan/admin/internal/common/convention/errorcode"

	"github.com/go-playground/validator/v10"
)

// FromBindError 区分输入错误和绑定调用错误，并保留原始原因。
func FromBindError(cause error) *ServiceError {
	if cause == nil {
		return nil
	}
	var invalidJSON *json.InvalidUnmarshalError
	var invalidValidation *validator.InvalidValidationError
	if errors.As(cause, &invalidJSON) || errors.As(cause, &invalidValidation) {
		return WrapServiceError(errorcode.ServiceError, cause)
	}
	if errors.Is(cause, io.EOF) {
		return WrapServiceError(errorcode.RequestBodyEmpty, cause)
	}
	var syntaxError *json.SyntaxError
	var typeError *json.UnmarshalTypeError
	var oversized *http.MaxBytesError
	if errors.Is(cause, io.ErrUnexpectedEOF) || errors.As(cause, &syntaxError) ||
		errors.As(cause, &typeError) || errors.As(cause, &oversized) {
		err := WrapServiceError(errorcode.RequestBodyInvalid, cause)
		if typeError != nil && typeError.Field != "" {
			err.Message = fieldLabel(typeError.Field) + "类型错误"
		}
		if oversized != nil {
			err.Message = "请求体超过大小限制"
		}
		return err
	}
	var validationErrors validator.ValidationErrors
	if errors.As(cause, &validationErrors) && len(validationErrors) > 0 {
		err := WrapServiceError(errorcode.ParameterInvalid, cause)
		err.Message = validationMessage(validationErrors[0])
		return err
	}
	var numberError *strconv.NumError
	if errors.As(cause, &numberError) {
		return WrapServiceError(errorcode.ParameterInvalid, cause)
	}
	// 包括传错绑定对象等未知异常，不能伪装成用户输入错误。
	return WrapServiceError(errorcode.ServiceError, cause)
}

func fieldLabel(field string) string {
	switch field {
	case "Username", "username":
		return "用户名"
	case "Password", "password":
		return "密码"
	case "RealName", "realName":
		return "姓名"
	case "Phone", "phone":
		return "手机号"
	case "Mail", "mail":
		return "邮箱"
	case "Token", "token":
		return "登录凭证"
	default:
		return "请求字段"
	}
}

func validationMessage(field validator.FieldError) string {
	label := fieldLabel(field.Field())
	switch field.Tag() {
	case "required":
		return label + "不能为空"
	case "min":
		return fmt.Sprintf("%s长度不能少于%s个字符", label, field.Param())
	case "max":
		return fmt.Sprintf("%s长度不能超过%s个字符", label, field.Param())
	case "len":
		return fmt.Sprintf("%s长度必须为%s个字符", label, field.Param())
	case "email":
		return "邮箱格式不正确"
	case "numeric":
		return label + "必须为数字"
	default:
		return label + "格式不正确"
	}
}
