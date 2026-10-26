package middleware

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

// RequestLogger 不记录 Query、Header、请求体，避免输出 token 和密码。
func RequestLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		started := time.Now()
		c.Next()
		log.Printf("请求完成 requestId=%s method=%s path=%q status=%d duration=%s",
			c.Writer.Header().Get("X-Request-ID"), c.Request.Method,
			c.Request.URL.Path, c.Writer.Status(), time.Since(started))
	}
}
