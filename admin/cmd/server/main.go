package main

import (
	"linkspan/admin/bootstrap"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	//创建Gin引擎
	r := gin.Default()

	//装配管理模块路由
	bootstrap.RegisterRoutes(r)

	//监听8001端口
	if err := r.Run(":8001"); err != nil {
		log.Fatal(err)
	}
}
