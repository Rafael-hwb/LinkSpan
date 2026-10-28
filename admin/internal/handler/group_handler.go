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
