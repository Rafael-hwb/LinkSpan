package router

import (
	"linkspan/admin/internal/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine, userHandler *handler.UserHandler) {
	adminGroup := r.Group("/api/short-link/admin/v1")

	adminGroup.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"module": "admin",
			"status": "ok",
		})
	})

	adminGroup.GET("user/:username", userHandler.GetUserByUsername)
}
