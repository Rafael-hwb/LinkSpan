package router

import (
	"linkspan/admin/internal/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine, userHandler *handler.UserHandler, requireLogin gin.HandlerFunc) {
	adminGroup := r.Group("/api/short-link/admin/v1")

	adminGroup.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"module": "admin",
			"status": "ok",
		})
	})

	adminGroup.GET("/user/:username", userHandler.GetUserByUsername)

	adminGroup.GET("/user/has-username", userHandler.HasUsername)

	adminGroup.POST("/user", userHandler.Register)

	adminGroup.PUT("/user", requireLogin, userHandler.Update)

	adminGroup.GET("/user/check-login", userHandler.CheckLogin)

	adminGroup.DELETE("/user/logout", userHandler.Logout)
}
