package service

import (
	"context"
	"database/sql"
	"fmt"
	"linkspan/admin/internal/cache"
	"linkspan/admin/internal/common/biz/usercontext"
	"linkspan/admin/internal/common/convention/errorcode"
	"linkspan/internal/common/convention/exception"
	"linkspan/admin/internal/dao/entity"
	"linkspan/admin/internal/dto/req"
	"linkspan/admin/internal/dto/resp"
	"linkspan/admin/internal/repository"
	"linkspan/admin/internal/toolkit"
	"log"
	"strings"
	"time"
	"unicode/utf8"
)

type GroupService struct {
	groupRepository       *repository.GroupRepository
	groupUniqueRepository *repository.GroupUniqueRepository
	groupLock             *cache.GroupCreateLock
	groupBloom            *cache.GroupBloom
	maxGroupCount         int
}

func NewGroupService(
	groupRepository *repository.GroupRepository,
	groupUniqueRepository *repository.GroupUniqueRepository,
	groupLock *cache.GroupCreateLock,
	groupBloom *cache.GroupBloom,
	maxGroupCount int,
) *GroupService {
	return &GroupService{
		groupRepository:       groupRepository,
		groupUniqueRepository: groupUniqueRepository,
		groupLock:             groupLock,
		groupBloom:            groupBloom,
		maxGroupCount:         maxGroupCount,
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

	if s.maxGroupCount <= 0 {
		return fmt.Errorf("分组数量上限配置无效")
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	token, acquired, err := s.groupLock.TryLock(ctx, username)
	if err != nil {
		return fmt.Errorf("获取分组创建锁失败: %w", err)
	}
	if !acquired {
		return exception.NewServiceError(errorcode.FlowLimitError)
	}
	defer s.releaseCreateLock(username, token)

	count, err := s.groupRepository.CountByUsername(ctx, username)
	if err != nil {
		return err
	}
	if count >= s.maxGroupCount {
		return exception.NewServiceError(errorcode.GroupLimitExceeded)
	}

	gid, err := s.createGroup(ctx, username, name)
	if err != nil {
		return err
	}
	s.recordGroupGID(gid)
	return nil
}

// createGroup 将全局 gid 登记和分组写入放在同一 MySQL 事务中。
func (s *GroupService) createGroup(ctx context.Context, username, name string) (string, error) {
	var gid string

	err := s.groupRepository.WithTransaction(ctx, func(tx *sql.Tx) error {
		var err error
		gid, err = s.InsertGroupInTx(ctx, tx, username, name)
		return err
	})
	if err != nil {
		return "", err
	}

	return gid, nil
}

// 空字符串表示过滤器预判已存在或唯一键冲突，调用方可以生成下一个候选。
func (s *GroupService) saveGroupUniqueReturnGID(ctx context.Context, tx *sql.Tx) (string, error) {
	gid, err := toolkit.GenerateGroupGID()
	if err != nil {
		return "", fmt.Errorf("生成分组标识失败: %w", err)
	}
	exists, err := s.groupBloom.Contains(ctx, gid)
	if err != nil {
		return "", err
	}
	if exists {
		return "", nil
	}
	reserved, err := s.groupUniqueRepository.Reserve(ctx, tx, gid)
	if err != nil {
		return "", err
	}
	if !reserved {
		return "", nil
	}
	return gid, nil
}

func (s *GroupService) releaseCreateLock(username, token string) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := s.groupLock.Unlock(ctx, username, token); err != nil {
		log.Printf("释放分组创建锁失败: %v", err)
	}
}

// 已提交后缓存失败只记录日志，避免客户端误以为创建失败而重复提交。
func (s *GroupService) recordGroupGID(gid string) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	if err := s.groupBloom.Add(ctx, gid); err != nil {
		log.Printf("分组已创建，但过滤器同步失败 gid=%s error=%v", gid, err)
	}
}

func (s *GroupService) List(ctx context.Context) ([]resp.GroupResp, error) {
	username, ok := usercontext.Username(ctx)
	if !ok {
		return nil, exception.NewServiceError(errorcode.LoginRequired)
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	groups, err := s.groupRepository.ListByUsername(ctx, username)
	if err != nil {
		return nil, err
	}

	response := make([]resp.GroupResp, 0, len(groups))

	for _, group := range groups {
		response = append(response, resp.GroupResp{
			GID:       group.GID,
			Name:      group.Name,
			SortOrder: group.SortOrder,
		})
	}

	return response, nil
}

func (s *GroupService) Delete(ctx context.Context, request req.GroupDeleteReq) error {
	username, ok := usercontext.Username(ctx)
	if !ok {
		return exception.NewServiceError(errorcode.LoginRequired)
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	return s.groupRepository.DeleteByGID(ctx, username, request.GID)
}

func (s *GroupService) UpdateName(ctx context.Context, request req.GroupUpdateReq) error {
	username, ok := usercontext.Username(ctx)
	if !ok {
		return exception.NewServiceError(errorcode.LoginRequired)
	}

	name := strings.TrimSpace(request.Name)

	if name == "" || utf8.RuneCountInString(name) > 64 {
		return exception.NewServiceError(errorcode.GroupNameInvalid)
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return s.groupRepository.UpdateName(ctx, username, request.GID, name)
}

func (s *GroupService) Sort(ctx context.Context, items []req.GroupSortReq) error {
	username, ok := usercontext.Username(ctx)
	if !ok {
		return exception.NewServiceError(errorcode.LoginRequired)
	}

	if len(items) == 0 {
		return exception.NewServiceError(errorcode.ParameterInvalid)
	}

	seen := make(map[string]struct{}, len(items))

	for _, item := range items {
		if item.GID == "" || len(item.GID) > 32 || item.SortOrder == nil {
			return exception.NewServiceError(errorcode.ParameterInvalid)
		}

		if _, exists := seen[item.GID]; exists {
			return exception.NewServiceError(errorcode.ParameterInvalid)
		}

		seen[item.GID] = struct{}{}
	}

	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	return s.groupRepository.WithTransaction(ctx, func(tx *sql.Tx) error {
		for _, item := range items {
			if err := s.groupRepository.UpdateSortOrder(ctx, tx, username, item.GID, *item.SortOrder); err != nil {
				return err
			}
		}

		return nil
	})
}

func (s *GroupService) RequireOwnedGroup(ctx context.Context, gid string) error {
	username, ok := usercontext.Username(ctx)
	if !ok {
		return exception.NewServiceError(errorcode.LoginRequired)
	}

	allowed, err := s.groupRepository.ExistsActiveByUsernameAndGID(ctx, username, gid)
	if err != nil {
		return err
	}

	if !allowed {
		return exception.NewServiceError(errorcode.GroupUnavailable)
	}

	return nil
}


func (s *GroupService) InsertGroupInTx(ctx context.Context, tx *sql.Tx, username, name string) (string, error) {
	for attempt := 0; attempt < 10; attempt++ {
		gid, err := s.saveGroupUniqueReturnGID(ctx, tx)
		if err != nil {
			return "", err
		}
		if gid == "" {
			continue
		}

		group := entity.GroupDO{
			GID:        gid,
			Name:       name,
			Username:   username,
			SortOrder:  0,
		}

		if err := s.groupRepository.Insert(ctx, tx, group); err != nil {
			return "", err
		}

		return gid, nil
	}

	return "", exception.NewServiceError(errorcode.GroupGIDGenerateFailed)
}
