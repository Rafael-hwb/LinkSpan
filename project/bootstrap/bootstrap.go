package bootstrap

import (
	"database/sql"

	"linkspan/project/internal/cache"
	"linkspan/project/internal/handler"
	"linkspan/project/internal/repository"
	"linkspan/project/internal/router"
	"linkspan/project/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
)

func RegisterRoutes(r *gin.Engine, db *sql.DB, redisClient *redis.Client) {
	links := repository.NewShortLinkRepository(db)
	routes := repository.NewShortLinkGotoRepository(db)
	gotoCache := cache.NewShortLinkGotoCache(redisClient)
	shortLinkService := service.NewShortLinkService(&links, routes, gotoCache)
	router.Register(r, handler.NewShortLinkHandler(shortLinkService))
}
