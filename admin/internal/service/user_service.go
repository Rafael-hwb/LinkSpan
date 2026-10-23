package service

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"linkspan/admin/internal/cache"
	"linkspan/admin/internal/common/convention/errorcode"
	"linkspan/admin/internal/common/convention/exception"
	"linkspan/admin/internal/common/serialize"
	"linkspan/admin/internal/dto/resp"
	"linkspan/admin/internal/repository"
)

type UserService struct {
	userRepository *repository.UserRepository
	userBloom      *cache.UserBloom
}

func NewUserService(userRepository *repository.UserRepository, userBloom *cache.UserBloom) *UserService {
	return &UserService{
		userRepository: userRepository,
		userBloom:      userBloom,
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
		Phone:    serialize.Phone(user.Phone),
		Mail:     user.Mail,
	}, nil
}

// true为可用， false为可能被占用
func (s *UserService) HasUsername(ctx context.Context, username string) (bool, error) {

	exists, err := s.userBloom.Contains(ctx, username)
	if err != nil {
		return false, err
	}

	return !exists, nil
}
