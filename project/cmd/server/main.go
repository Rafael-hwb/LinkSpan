package main

import (
	"log"

	adminbootstrap "linkspan/admin/bootstrap"
	"linkspan/internal/middleware"
	"linkspan/project/bootstrap"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := adminbootstrap.OpenDatabase()
	if err != nil {
		log.Fatalf("连接 MySQL 失败: %v", err)
	}
	defer db.Close()

	r := gin.New()
	r.Use(middleware.RequestLogger(), middleware.ErrorHandler())

	bootstrap.RegisterRoutes(r, db)

	if err := r.Run("127.0.0.1:8002"); err != nil {
		log.Fatalf("启动短链接核心服务失败: %v", err)
	}
}
