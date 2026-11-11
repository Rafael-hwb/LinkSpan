package exception_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"testing"

	"linkspan/admin/internal/dto/req"
	"linkspan/internal/common/convention/errorcode"
	"linkspan/internal/common/convention/exception"

	"github.com/gin-gonic/gin/binding"
)

func TestFromBindError(t *testing.T) {
	var destination req.UserLoginReq
	invalidPointer := json.Unmarshal([]byte(`{}`), destination)
	invalidType := json.Unmarshal([]byte(`{"username":123}`), &destination)
	invalidJSON := json.Unmarshal([]byte(`{"username":}`), &destination)
	missingField := binding.Validator.ValidateStruct(&req.UserLoginReq{Username: "test3"})
	invalidMail := binding.Validator.ValidateStruct(&req.UserRegisterReq{
		Username: "test3", Password: "TestPass123!", RealName: "测试", Phone: "13800000000", Mail: "invalid",
	})
	tests := []struct {
		name    string
		err     error
		code    string
		message string
	}{
		{"empty", io.EOF, "A000004", "请求体不能为空"},
		{"truncated", io.ErrUnexpectedEOF, "A000002", "请求体格式错误"},
		{"syntax", invalidJSON, "A000002", "请求体格式错误"},
		{"type", invalidType, "A000002", "用户名类型错误"},
		{"required", missingField, "A000003", "密码不能为空"},
		{"email", invalidMail, "A000003", "邮箱格式不正确"},
		{"non_pointer_is_program_error", invalidPointer, "B000001", errorcode.ServiceError.Message},
		{"unknown_is_program_error", errors.New("binding configuration problem"), "B000001", errorcode.ServiceError.Message},
		{"oversized", &http.MaxBytesError{Limit: 1024}, "A000002", "请求体超过大小限制"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.err == nil {
				t.Fatal("fixture did not produce an error")
			}
			wrapped := fmt.Errorf("bind: %w", tt.err)
			got := exception.FromBindError(wrapped)
			if got.Code != tt.code || got.Message != tt.message {
				t.Fatalf("got %s %q, want %s %q", got.Code, got.Message, tt.code, tt.message)
			}
			if !errors.Is(got, wrapped) {
				t.Fatal("original cause was lost")
			}
		})
	}
	if exception.FromBindError(nil) != nil {
		t.Fatal("nil error should stay nil")
	}
}
