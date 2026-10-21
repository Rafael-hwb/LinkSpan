package middleware

import (
	"errors"
	"linkspan/admin/internal/common/convention/errorcode"
	"linkspan/admin/internal/common/convention/exception"
	"linkspan/admin/internal/common/convention/result"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			//处理请求中意外产生的panic
			if recovered := recover(); recovered != nil {
				log.Printf(
					"请求发生panic: %v\n%s",
					recovered,
					debug.Stack(),
				)

				_ = c.Error(errors.New("请求处理发生 panic"))
				c.Abort()
			}
			
			//没有错误，交给正常流程
			if len(c.Errors) == 0 {
				return 
			}
			
			//检查和记录错误
			err := c.Errors.Last().Err

			log.Printf(
				"请求失败 method=%s path=%s error=%v",
				c.Request.Method,
				c.Request.URL.Path,
				err,
			)

			// 已经发送响应时，不再写入第二份响应。
			if c.Writer.Written() {
				return
			}

			//设置默认值
			code := errorcode.ServiceError.Code
			message := errorcode.ServiceError.Message

			// 识别已知业务错误
			var ServiceError *exception.ServiceError
			if errors.As(err, &ServiceError) {
				code = ServiceError.Code
				message = ServiceError.Message
			}

			c.AbortWithStatusJSON(
				http.StatusOK,
				result.Failure(code, message),
			)
		}()
		c.Next()
	}
}