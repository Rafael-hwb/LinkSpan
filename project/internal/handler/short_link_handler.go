package handler

import (
	"errors"
	"net/http"
	"strings"

	"linkspan/internal/common/convention/exception"
	"linkspan/internal/common/convention/result"
	"linkspan/project/internal/dto/req"
	"linkspan/project/internal/service"

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

	response, err := h.service.Create(c.Request.Context(), request)
	if err != nil {
		_ = c.Error(err)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, result.Success(response))
}

func (h *ShortLinkHandler) Page(c *gin.Context) {
	var request req.ShortLinkPageReq
	if err := c.ShouldBindQuery(&request); err != nil {
		_ = c.Error(exception.FromBindError(err))
		c.Abort()
		return
	}

	response, err := h.service.Page(c.Request.Context(), request)
	if err != nil {
		_ = c.Error(err)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, result.Success(response))
}

func (h *ShortLinkHandler) Update(c *gin.Context) {
	var request req.ShortLinkUpdateReq
	if err := c.ShouldBindJSON(&request); err != nil {
		_ = c.Error(exception.FromBindError(err))
		c.Abort()
		return
	}

	if err := h.service.Update(c.Request.Context(), request); err != nil {
		_ = c.Error(err)
		c.Abort()
		return
	}
	c.JSON(http.StatusOK, result.Success[any](nil))
}

func (h *ShortLinkHandler) RestoreURL(c *gin.Context) {
	if c.Request.Method != http.MethodGet {
		c.Status(http.StatusNotFound)
		return
	}

	path := strings.TrimPrefix(c.Request.URL.Path, "/")
	if path == "" || strings.Contains(path, "/") {
		c.Status(http.StatusNotFound)
		return
	}

	fullShortURL := c.Request.Host + "/" + path
	origin, err := h.service.RestoreURL(c.Request.Context(), fullShortURL)
	if err != nil {
		if errors.Is(err, service.ErrShortLinkNotFound) {
			c.Status(http.StatusNotFound)
			return
		}
		_ = c.Error(err)
		return
	}
	c.Redirect(http.StatusFound, origin)
}
