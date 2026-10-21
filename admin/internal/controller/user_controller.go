package controller

import (
	"linkspan/admin/internal/common/convention/result"
	"linkspan/admin/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService *service.UserService
}

func NewUserController(userService *service.UserService) *UserController {
	return &UserController{
		userService: userService,
	}
}

func (u *UserController) GetUserByUsername(c *gin.Context) {
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