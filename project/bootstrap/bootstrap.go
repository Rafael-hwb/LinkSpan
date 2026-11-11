package bootstrap

import (
	"database/sql"

	"linkspan/project/internal/handler"
	"linkspan/project/internal/repository"
	"linkspan/project/internal/router"
	"linkspan/project/internal/service"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, db *sql.DB) {
	links := repository.NewShortLinkRepository(db)
	routes := repository.NewShortLinkGotoRepository(db)
	shortLinkService := service.NewShortLinkService(&links, routes)
	router.Register(r, handler.NewShortLinkHandler(shortLinkService))
}
