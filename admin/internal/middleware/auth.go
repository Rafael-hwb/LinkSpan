package middleware

import (
	"linkspan/admin/internal/cache"
	"linkspan/admin/internal/common/biz/usercontext"
	"linkspan/admin/internal/common/convention/errorcode"
	"linkspan/admin/internal/common/convention/exception"

	"github.com/gin-gonic/gin"
)

func RequireLogin(session *cache.UserSession) gin.HandlerFunc {
	return func(c *gin.Context) {
		username := c.GetHeader("username")
		token := c.GetHeader("token")

		if username == "" || token == "" {
			rejectLogin(c, errorcode.LoginRequired)
			return
		}

		valid, err := session.Exists(c.Request.Context(), username, token)

		if err != nil {
			_ = c.Error(err)
			c.Abort()
			return
		}

		if !valid {
			rejectLogin(c, errorcode.SessionInvalid)
			return
		}

		ctx := usercontext.WithUsername(
			c.Request.Context(),
			username,
		)

		c.Request = c.Request.WithContext(ctx)

		c.Next()
	}
}

func rejectLogin(c *gin.Context, code errorcode.ErrorCode) {
	_ = c.Error(exception.NewServiceError(code))
	c.Abort()
}
