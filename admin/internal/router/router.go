package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	adminGroup := r.Group("/api/shortlink/admin/vi")

	adminGroup.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"module": "admin",
			"status": "ok",
		})
	})
}
