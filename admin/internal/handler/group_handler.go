package handler

import (
	"net/http"

	"linkspan/admin/internal/common/convention/exception"
	"linkspan/admin/internal/common/convention/result"
	"linkspan/admin/internal/dto/req"
	"linkspan/admin/internal/service"

	"github.com/gin-gonic/gin"
)

type GroupHandler struct {
	groupService *service.GroupService
}

func NewGroupHandler(groupService *service.GroupService) *GroupHandler {
	return &GroupHandler{groupService: groupService}
}

func (h *GroupHandler) Save(c *gin.Context) {
	var request req.GroupSaveReq
	if err := c.ShouldBindJSON(&request); err != nil {
		_ = c.Error(exception.FromBindError(err))
		c.Abort()
		return
	}
	if err := h.groupService.Save(c.Request.Context(), request); err != nil {
		_ = c.Error(err)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, result.Success[any](nil))
}

func (h *GroupHandler) List(c *gin.Context) {
	response, err := h.groupService.List(c.Request.Context())
	if err != nil {
		_ = c.Error(err)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, result.Success(response))
}

func (h *GroupHandler) Delete(c *gin.Context) {
	var request req.GroupDeleteReq
	if err := c.ShouldBindQuery(&request); err != nil {
		_ = c.Error(exception.FromBindError(err))
		c.Abort()
		return
	}
	if err := h.groupService.Delete(c.Request.Context(), request); err != nil {
		_ = c.Error(err)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, result.Success[any](nil))
}

func (h *GroupHandler) UpdateName(c *gin.Context) {
	var request req.GroupUpdateReq

	if err := c.ShouldBindJSON(&request); err != nil {
		_ = c.Error(exception.FromBindError(err))
		c.Abort()
		return
	}

	if err := h.groupService.UpdateName(c.Request.Context(), request); err != nil {
		_ = c.Error(err)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, result.Success[any](nil))
}

func (h *GroupHandler) Sort(c *gin.Context) {
	var items []req.GroupSortReq

	if err := c.ShouldBindJSON(&items); err != nil {
		_ = c.Error(exception.FromBindError(err))
		c.Abort()
		return
	}

	if err := h.groupService.Sort(c.Request.Context(), items); err != nil {
		_ = c.Error(err)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, result.Success[any](nil))
}
