package service

import (
	"context"
	"fmt"
	"linkspan/admin/internal/common/biz/usercontext"
	"linkspan/admin/internal/common/convention/errorcode"
	"linkspan/admin/internal/common/convention/exception"
	"linkspan/admin/internal/dto/req"
	"linkspan/admin/internal/repository"
	"strings"
	"time"
	"unicode/utf8"
)

type GroupService struct {
	groupRepository *repository.GroupRepository
	maxGroupInt     int
}

func (s *GroupService) NewGroupService(groupRepository *repository.GroupRepository, maxGroupInt int) *GroupService {
	return &GroupService{
		groupRepository: groupRepository,
		maxGroupInt: maxGroupInt,
	}
}

func (s *GroupService) Save(ctx context.Context, request req.GroupSaveReq) error {
	username, ok := usercontext.Username(ctx)

	if !ok {
		return exception.NewServiceError(errorcode.LoginRequired)
	}

	name := strings.TrimSpace(request.Name)

	if name == "" || utf8.RuneCountInString(name) > 64 {
		return exception.NewServiceError(errorcode.GroupNameInvalid)
	}

	if s.maxGroupInt <= 0 {
		return fmt.Errorf("分组数量上限配置无效")
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return 
}

func generateGroupGID() (string, error) {
	
}