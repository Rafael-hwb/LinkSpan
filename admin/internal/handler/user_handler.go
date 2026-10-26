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

func (h *UserHandler) Register(c *gin.Context) {
	var request req.UserRegisterReq

	if err := c.ShouldBindJSON(&request); err != nil {
		_ = c.Error(exception.NewServiceError(errorcode.ClientError))
		c.Abort()
		return
	}

	if err := h.userService.Register(c.Request.Context(), request); err != nil {
		_ = c.Error(err)
		c.Abort()
		return
	}

	c.JSON(
		http.StatusOK,
		result.Success[any](nil),
	)
}

func (h *UserHandler) Update(c *gin.Context) {
	var request *req.UserUpdateReq

	if err := c.ShouldBindJSON(&request); err != nil {
		_ = c.Error(exception.NewServiceError(errorcode.ClientError))
		c.Abort()
		return
	}

	if err := h.userService.Update(c.Request.Context(), *request); err != nil {
		_ = c.Error(err)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, result.Success[any](nil))
}

func (h *UserHandler) Login(c *gin.Context) {
	var request req.UserLoginReq

	if err := c.ShouldBindJSON(request); err != nil {
		_ = c.Error(exception.NewServiceError(errorcode.ClientError))
		c.Abort()
		return
	}

	response, err := h.userService.Login(c.Request.Context(), request)

	if err != nil {
		_ = c.Error(err)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, result.Success(response))

}

func (h *UserHandler) CheckLogin(c *gin.Context) {
	var request req.UserSessionReq

	if err := c.ShouldBindJSON(&request); err != nil {
		_ = c.Error(
			exception.NewServiceError(errorcode.ClientError),
		)
		c.Abort()
		return
	}

	valid, err := h.userService.CheckLogin(c.Request.Context(), request.Username, request.Token)

	if err != nil {
		_ = c.Error(err)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, result.Success(valid))
}

func (h *UserHandler) Logout(c *gin.Context) {
	var request req.UserSessionReq

	if err := c.ShouldBindJSON(&request); err != nil {
		_ = c.Error(exception.NewServiceError(errorcode.ClientError))
		c.Abort()
		return
	}

	if err := h.userService.Logout(c.Request.Context(), request.Username, request.Token); err != nil {
		_ = c.Error(err)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, result.Success[any](nil))
}
