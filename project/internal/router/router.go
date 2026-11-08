package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine) {
	projectGroup := r.Group("/api/short-link/v1")

	projectGroup.GET("health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"module": "project",
			"status": "ok",
		})
	})
}