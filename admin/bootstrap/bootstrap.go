package bootstrap

import (
	"context"
	"database/sql"
	"fmt"
	"linkspan/admin/internal/cache"
	"linkspan/admin/internal/handler"
	"linkspan/admin/internal/middleware"
	"linkspan/admin/internal/remote"
	"linkspan/admin/internal/repository"
	"linkspan/admin/internal/router"
	"linkspan/admin/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

// 向已有的Gin引擎注册管理模块路由
func RegisterRoutes(ctx context.Context, r *gin.Engine,
	db *sql.DB, redisClient *redis.Client) error {
	fieldCipher, err := OpenFieldCipher()
	if err != nil {
		return err
	}

	userRepository := repository.NewUserRepository(db, fieldCipher)
	userBloom := cache.NewUserBloom(redisClient)
	userLock := cache.NewUserRegisterLock(redisClient)
	userSession := cache.NewUserSession(redisClient)

	if err := userBloom.Prepare(ctx); err != nil {
		return err
	}

	if err := userRepository.ForEachUsername(
		ctx,
		func(username string) error {
			return userBloom.Add(ctx, username)
		},
	); err != nil {
		return fmt.Errorf("导入已有用户名失败: %w", err)
	}

	if err := userBloom.MarkReady(ctx); err != nil {
		return fmt.Errorf("标记过滤器就绪失败: %w", err)
	}

	userService := service.NewUserService(userRepository, userBloom, userLock, userSession)
	userHandler := handler.NewUserHandler(userService)

	groupRepository := repository.NewGroupRepository(db)
	groupUniqueRepository := repository.NewGroupUniqueRepository(db)
	groupLock := cache.NewGroupCreateLock(redisClient)
	groupBloom := cache.NewGroupBloom(redisClient)

	if err := groupBloom.Prepare(ctx); err != nil {
		return err
	}
	if err := groupUniqueRepository.ForEachGID(ctx, func(gid string) error {
		return groupBloom.Add(ctx, gid)
	}); err != nil {
		return fmt.Errorf("导入已有分组标识失败: %w", err)
	}
	if err := groupBloom.MarkReady(ctx); err != nil {
		return err
	}

	groupService := service.NewGroupService(groupRepository, groupUniqueRepository, groupLock, groupBloom, 20)
	groupHandler := handler.NewGroupHandler(groupService)
	shortLinkService := service.NewShortLinkService(groupService, remote.NewShortLinkClient())
	shortLinkHandler := handler.NewShortLinkHandler(shortLinkService)
	router.Register(r, userHandler, groupHandler, shortLinkHandler, middleware.RequireLogin(userSession))

	return nil
}
