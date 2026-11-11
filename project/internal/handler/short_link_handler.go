package handler

import (
	"errors"
	"log"
	"net/http"

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
		writeCreateFailure(c, "A000003", "请检查创建短链接的请求参数")
		return
	}

	response, err := h.service.Create(c.Request.Context(), request)
	switch {
	case err == nil:
		c.JSON(http.StatusOK, gin.H{
			"code": "0", "message": "", "data": response, "requestId": nil, "success": true,
		})
	case errors.Is(err, service.ErrInvalidCreateRequest):
		writeCreateFailure(c, "A000003", err.Error())
	default:
		log.Printf("创建短链接失败: %v", err)
		writeCreateFailure(c, "B000001", "系统执行错误")
	}
}

func writeCreateFailure(c *gin.Context, code, message string) {
	c.JSON(http.StatusOK, gin.H{
		"code": code, "message": message, "data": nil, "requestId": nil, "success": false,
	})
}

func (h *ShortLinkHandler) Page(c *gin.Context) {
	var request req.ShortLinkPageReq
	if err := c.ShouldBindQuery(&request); err != nil {
		writeCreateFailure(c, "A000003", "请检查分页参数")
		return
	}

	response, err := h.service.Page(c.Request.Context(), request)
	if err != nil {
		log.Printf("分页查询短链接失败: %v", err)
		writeCreateFailure(c, "B000001", "系统执行错误")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code": "0", "message": "", "data": response,
		"requestId": nil, "success": true,
	})
}

func (h *ShortLinkHandler) Update(c *gin.Context) {
	var request req.ShortLinkUpdateReq
	if err := c.ShouldBindJSON(&request); err != nil {
		writeCreateFailure(c, "A000003", "请检查创建短链接的请求参数")
		return
	}

	err := h.service.Update(c.Request.Context(), request)
	switch {
	case err == nil:
		c.JSON(http.StatusOK, gin.H{
			"code": "0", "message": "", "data": nil, "requestId": nil, "success": true,
		})
	case errors.Is(err, service.ErrInvalidCreateRequest):
		writeCreateFailure(c, "A000003", err.Error())
	default:
		log.Printf("创建短链接失败: %v", err)
		writeCreateFailure(c, "B000001", "系统执行错误")
	}
}