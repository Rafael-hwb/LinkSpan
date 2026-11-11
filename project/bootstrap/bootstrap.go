package bootstrap

import (
	"context"
	"database/sql"
	"fmt"

	"linkspan/project/internal/cache"
	"linkspan/project/internal/handler"
	"linkspan/project/internal/repository"
	"linkspan/project/internal/router"
	"linkspan/project/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func RegisterRoutes(ctx context.Context, r *gin.Engine, db *sql.DB, redisClient *redis.Client) error {
	links := repository.NewShortLinkRepository(db)
	routes := repository.NewShortLinkGotoRepository(db)
	gotoCache := cache.NewShortLinkGotoCache(redisClient)
	bloom := cache.NewShortLinkBloom(redisClient)
	if err := bloom.Prepare(ctx); err != nil {
		return err
	}
	if err := routes.ForEachFullShortURL(ctx, func(fullShortURL string) error {
		return bloom.Add(ctx, fullShortURL)
	}); err != nil {
		return fmt.Errorf("导入已有短链接失败: %w", err)
	}
	if err := bloom.MarkReady(ctx); err != nil {
		return err
	}
	shortLinkService := service.NewShortLinkService(&links, routes, gotoCache, bloom)
	router.Register(r, handler.NewShortLinkHandler(shortLinkService))
	return nil
}
