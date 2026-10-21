package router

import (
	"linkspan/admin/internal/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine, UserController *controller.UserController) {
	adminGroup := r.Group("/api/short-link/admin/v1")

	adminGroup.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"module": "admin",
			"status": "ok",
		})
	})

	adminGroup.GET("user/:username", UserController.GetUserByUsername)
}
