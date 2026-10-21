package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"linkspan/admin/internal/common/convention/errorcode"
	"linkspan/admin/internal/common/convention/exception"
	"linkspan/admin/internal/dao/mapper"
	"linkspan/admin/internal/dto/resp"
)

type UserService struct {
	userMapper *mapper.UserMapper
}

func NewUserService(userMapper *mapper.UserMapper) *UserService {
	return &UserService{
		userMapper: userMapper,
	}
}

func (s *UserService) GetUserByUsername(ctx context.Context, username string) (*resp.UserResp, error) {
	user, err := s.userMapper.SelectByUsername(ctx, username)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, exception.NewServiceError(errorcode.UserNull)
	}

	if err != nil {
		return nil, fmt.Errorf("查询用户失败: %w", err)
	}

	return &resp.UserResp{
		ID:       user.ID,
		Username: user.Username,
		RealName: user.RealName,
		Phone:    user.Phone,
		Mail:     user.Mail,
	}, nil
}
