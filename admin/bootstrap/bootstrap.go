package bootstrap

import (
	"database/sql"
	"linkspan/admin/internal/controller"
	"linkspan/admin/internal/dao/mapper"
	"linkspan/admin/internal/router"
	"linkspan/admin/internal/service"

	"github.com/gin-gonic/gin"
)

//向已有的Gin引擎注册管理模块路由
func RegisterRoutes(r *gin.Engine, db *sql.DB) {
	userMapper := mapper.NewUserMapper(db)
	userService := service.NewUserService(userMapper)
	userController := controller.NewUserController(userService)
	router.Register(r, userController)
}
