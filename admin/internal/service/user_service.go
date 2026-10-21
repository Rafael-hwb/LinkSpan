package service

import (
	"context"
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
	
}