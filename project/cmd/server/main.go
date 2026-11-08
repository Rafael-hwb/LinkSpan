package server

import (
	"linkspan/project/bootstrap"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	bootstrap.RegisterRoutes(r)

	if err := r.Run(":8002"); err != nil {
		log.Fatalf("启动短链接核心服务失败: %v", err)
	}
}