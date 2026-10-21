package main

import (
	"linkspan/admin/bootstrap"
	"linkspan/admin/internal/middleware"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	//建立数据库连接
	db, err := bootstrap.OpenDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	//创建Gin引擎
	r := gin.New()

	//装配中间件
	r.Use(gin.Logger())
	r.Use(middleware.ErrorHandler())

	//装配管理模块路由
	bootstrap.RegisterRoutes(r, db)

	//监听8001端口
	if err := r.Run(":8001"); err != nil {
		log.Fatal(err)
	}
}
