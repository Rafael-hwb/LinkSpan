package router

import (
	"net/http"

	"linkspan/project/internal/handler"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine, shortLinkHandler *handler.ShortLinkHandler) {
	projectGroup := r.Group("/api/short-link/v1")

	projectGroup.GET("health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"module": "project",
			"status": "ok",
		})
	})
	projectGroup.POST("/create", shortLinkHandler.Create)

	projectGroup.GET("/page", shortLinkHandler.Page)
	projectGroup.POST("/update", shortLinkHandler.Update)
}
