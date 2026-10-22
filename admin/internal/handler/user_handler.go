package handler

import (
	"linkspan/admin/internal/common/convention/result"
	"linkspan/admin/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService *service.UserService
}

func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

func (u *UserHandler) GetUserByUsername(c *gin.Context) {
	username := c.Param("username")

	user, err := u.userService.GetUserByUsername(
		c.Request.Context(),
		username,
	)

	if err != nil {
		_ = c.Error(err)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, result.Success(user))
}
