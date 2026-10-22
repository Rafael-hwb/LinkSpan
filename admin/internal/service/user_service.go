package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"linkspan/admin/internal/common/convention/errorcode"
	"linkspan/admin/internal/common/convention/exception"
	"linkspan/admin/internal/dto/resp"
	"linkspan/admin/internal/repository"
)

type UserService struct {
	userRepository *repository.UserRepository
}

func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

func (s *UserService) GetUserByUsername(ctx context.Context, username string) (*resp.UserResp, error) {
	user, err := s.userRepository.SelectByUsername(ctx, username)
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
