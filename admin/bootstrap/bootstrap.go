package bootstrap

import (
	"linkspan/admin/internal/router"

	"github.com/gin-gonic/gin"
)

//向已有的Gin引擎注册管理模块路由
func RegisterRoutes(r *gin.Engine) {
	router.Register(r)
}