package handler

import (
	"net/http"

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
		return
	}

	response, err := h.service.Create(c.Request.Context(), request)
	if err != nil {
		_ = c.Error(err)
		return
	}
	writeSuccess(c, response)
}

func (h *ShortLinkHandler) Page(c *gin.Context) {
	var request req.ShortLinkPageReq
	if err := c.ShouldBindQuery(&request); err != nil {
		_ = c.Error(exception.FromBindError(err))
		return
	}

	response, err := h.service.Page(c.Request.Context(), request)
	if err != nil {
		_ = c.Error(err)
		return
	}
	writeSuccess(c, response)
}

func (h *ShortLinkHandler) Update(c *gin.Context) {
	var request req.ShortLinkUpdateReq
	if err := c.ShouldBindJSON(&request); err != nil {
		_ = c.Error(exception.FromBindError(err))
		return
	}

	if err := h.service.Update(c.Request.Context(), request); err != nil {
		_ = c.Error(err)
		return
	}
	writeSuccess(c, nil)
}

func writeSuccess(c *gin.Context, data any) {
	response := result.Success(data)
	requestID := c.Writer.Header().Get("X-Request-ID")
	response.RequestID = &requestID
	c.JSON(http.StatusOK, response)
}
