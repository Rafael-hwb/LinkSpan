package handler

import (
	"linkspan/admin/internal/common/convention/errorcode"
	"linkspan/admin/internal/common/convention/exception"
	"linkspan/admin/internal/common/convention/result"
	"linkspan/admin/internal/dto/req"
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

func (h *UserHandler) GetUserByUsername(c *gin.Context) {
	username := c.Param("username")

	user, err := h.userService.GetUserByUsername(
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


func (h *UserHandler) HasUsername(c *gin.Context) {
	var Request req.UserHasUsernameReq

	if err := c.ShouldBindQuery(&Request); err != nil {
		_ = c.Error(exception.NewServiceError(errorcode.ClientError))
		c.Abort()
		return
	}

	avaliable, err := h.userService.HasUsername(c.Request.Context(), Request.Username)
	
	if err != nil {
		c.Error(err)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, result.Success(avaliable))
}