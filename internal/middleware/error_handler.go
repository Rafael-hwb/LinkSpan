package middleware

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"log"
	"net/http"
	"runtime/debug"

	"linkspan/internal/common/convention/errorcode"
	"linkspan/internal/common/convention/exception"
	"linkspan/internal/common/convention/result"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestID := rand.Text()
		c.Header("X-Request-ID", requestID)

		defer func() {
			if recovered := recover(); recovered != nil {
				// 不输出可能携带个人信息的 panic 值，保留类型及堆栈。
				log.Printf("请求 panic requestId=%s type=%T\n%s", requestID, recovered, debug.Stack())
				_ = c.Error(fmt.Errorf("请求处理发生 panic (%T)", recovered))
				c.Abort()
			}
			if len(c.Errors) == 0 {
				return
			}
			err := c.Errors.Last().Err
			cause := err
			for errors.Unwrap(cause) != nil {
				cause = errors.Unwrap(cause)
			}
			log.Printf("请求失败 requestId=%s method=%s path=%q error=%v cause=%v",
				requestID, c.Request.Method, c.Request.URL.Path, err, cause)

			if c.Writer.Written() {
				return
			}
			code := errorcode.ServiceError.Code
			message := errorcode.ServiceError.Message
			var serviceErr *exception.ServiceError
			if errors.As(err, &serviceErr) {
				code = serviceErr.Code
				message = serviceErr.Message
			} else if errors.Is(err, context.DeadlineExceeded) {
				code = errorcode.ServiceTimeout.Code
				message = errorcode.ServiceTimeout.Message
			}
			response := result.Failure(code, message)
			response.RequestID = &requestID
			c.AbortWithStatusJSON(http.StatusOK, response)
		}()
		c.Next()
	}
}
