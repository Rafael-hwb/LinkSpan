package main

import (
	"context"
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

	redisClient, err := adminbootstrap.OpenRedis()
	if err != nil {
		log.Fatalf("连接 Redis 失败: %v", err)
	}
	defer redisClient.Close()

	r := gin.New()
	r.Use(middleware.RequestLogger(), middleware.ErrorHandler())

	if err := bootstrap.RegisterRoutes(context.Background(), r, db, redisClient); err != nil {
		log.Fatalf("初始化短链接路由失败: %v", err)
	}

	if err := r.Run("127.0.0.1:8002"); err != nil {
		log.Fatalf("启动短链接核心服务失败: %v", err)
	}
}
