package service

import (
	"context"
	"errors"
	"strings"
	"testing"

	"linkspan/admin/internal/common/biz/usercontext"
	"linkspan/admin/internal/common/convention/exception"
	"linkspan/admin/internal/dto/req"
)

func TestUserServiceRejectsUnauthorizedAndInvalidRequests(t *testing.T) {
	s := &UserService{} // 验证这些分支不依赖数据库或 Redis。
	ctx := usercontext.WithUsername(context.Background(), "test3")
	tests := []struct {
		name, code string
		run        func() error
	}{
		{"missing_identity", "A000400", func() error { return s.Update(context.Background(), req.UserUpdateReq{Username: "test3"}) }},
		{"other_user", "A000403", func() error { return s.Update(ctx, req.UserUpdateReq{Username: "other"}) }},
		{"no_fields", "A000005", func() error { return s.Update(ctx, req.UserUpdateReq{Username: "test3"}) }},
		{"register_username", "A000110", func() error { return s.Register(ctx, req.UserRegisterReq{Username: "INVALID"}) }},
		{"register_password_bytes", "A000120", func() error {
			return s.Register(ctx, req.UserRegisterReq{Username: "test3", Password: strings.Repeat("密", 25)})
		}},
		{"login_username", "A000110", func() error { _, err := s.Login(ctx, req.UserLoginReq{Username: "INVALID"}); return err }},
		{"login_password_bytes", "A000120", func() error {
			_, err := s.Login(ctx, req.UserLoginReq{Username: "test3", Password: strings.Repeat("密", 25)})
			return err
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var businessErr *exception.ServiceError
			if err := tt.run(); !errors.As(err, &businessErr) || businessErr.Code != tt.code {
				t.Fatalf("got %v, want code %s", err, tt.code)
			}
		})
	}
	cause := errors.New("credential diagnostic")
	err := loginFailedError(cause)
	if err.Code != "A000402" || !errors.Is(err, cause) {
		t.Fatal("credential error lost its classification or cause")
	}
}
