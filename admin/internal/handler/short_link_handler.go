package handler

import (
	"net/http"

	"linkspan/admin/internal/dto/req"
	"linkspan/admin/internal/service"
	"linkspan/internal/common/convention/exception"

	"github.com/gin-gonic/gin"
)

type ShortLinkHandler struct {
	service *service.ShortLinkService
}

func NewShortLinkHandler(s *service.ShortLinkService) *ShortLinkHandler {
	return &ShortLinkHandler{service: s}
}

func (h *ShortLinkHandler) Create(c *gin.Context) {
	var request req.ShortLinkCreateReq
	if err := c.ShouldBindJSON(&request); err != nil {
		_ = c.Error(exception.FromBindError(err))
		c.Abort()
		return
	}
	reply, err := h.service.Create(c.Request.Context(), request)
	if err != nil {
		_ = c.Error(err)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, reply)
}

func (h *ShortLinkHandler) Page(c *gin.Context) {
	var request req.ShortLinkPageReq
	if err := c.ShouldBindQuery(&request); err != nil {
		_ = c.Error(exception.FromBindError(err))
		c.Abort()
		return
	}
	reply, err := h.service.Page(c.Request.Context(), request)
	if err != nil {
		_ = c.Error(err)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, reply)
}

func (h *ShortLinkHandler) Update(c *gin.Context) {
	var request req.ShortLinkUpdateReq
	if err := c.ShouldBindJSON(&request); err != nil {
		_ = c.Error(exception.FromBindError(err))
		c.Abort()
		return
	}

	reply, err := h.service.Update(c.Request.Context(), request)
	if err != nil {
		_ = c.Error(err)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, reply)
}
