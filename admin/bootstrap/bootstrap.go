package bootstrap

import (
	"database/sql"
	"linkspan/admin/internal/handler"
	"linkspan/admin/internal/repository"
	"linkspan/admin/internal/router"
	"linkspan/admin/internal/service"

	"github.com/gin-gonic/gin"
)

// 向已有的Gin引擎注册管理模块路由
func RegisterRoutes(r *gin.Engine, db *sql.DB) {
	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userHandler := handler.NewUserHandler(userService)
	router.Register(r, userHandler)
}
